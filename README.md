# textWeb

textWeb is a small project that aims to help breakdown websites back into basic HTML.

## Why does this exist?
textWeb was originally created for a family member who is currently living in a more rural location, with only
one cell tower within their living vicinity, and no other form of internet.

This individual has a smartphone, but found it incredibly difficult to get pages to load when other visitors were
in the area due to network congestion (assumably at the cell). To help remedy that, we perform some stripping using
[Go-Readability](github.com/go-shiori/go-readability) which is in turn a Golang port of Mozilla's [Readability.js](https://github.com/mozilla/readability).

Performing the HTML stripping this way ensures that the "bloated" page is quickly accessed by a remote server, and then
parsed down to basic HTML, _before_ sending it along to the constricted mobile device.

This tool also has the added benefit of stripping pages down in a readable format for older devices, i.e. my
Macintosh Plus. Though, some additional work to remove more js and other element compatibility is still required
to actually achieve full functionality.

## How to use it
Easily enough, clone the repo, `cd` into it, and run `make` (assuming you have Go installed) and you should get a nice
binary in `build/` that you can run.


## How to use it on a server
If you want to run this on a server, you _can_ run it without a reverse proxy, however, for my running implementation,
I have it reverse-proxied behind Nginx. I'll let Nginx run on the privileged ports. I don't run this with SSL (though you can)
to allow compatibility with older devices that don't support SSL. Since readability requests don't send inherently
"secure" content, I don't see this as a big deal, but you can definitely use SSL on Nginx, and YMMV.

To set this up as a service, simply copy the binary (you may need to run `make build-all` to get the binary for the arch
you need) along with the `templates/` to a dir on your server (i.e. /var/lib/textweb).

For security purposes, you might want to run this as its own user, so running something like 
```
sudo useradd textweb -m -d /var/lib/textweb
```
is probably not a bad idea either. From there, create a systemD unit (see below) and then configure your nginx
reverse proxy (see below also). Then start your service, and get going.

**SystemD Example**
```
[Unit]
Description=TextWeb
After=syslog.target
After=network.target

[Service]
RestartSec=2s
Type=simple
User=textweb
Group=textweb
WorkingDirectory=/var/lib/textweb

ExecStart=/usr/local/bin/textWeb
Restart=always
Environment=GIN_MODE=release

[Install]
WantedBy=multi-user.target
```

**Nginx Config Example**
```
server {
    listen 80;
    server_name textweb.somesite.org;

    location / {
       proxy_pass http://127.0.0.1:8080;
    }
}
```

#### the commands
```
# Create textWeb User
sudo useradd textweb -m -d /var/lib/textweb

# Copy Binary
sudo mv reader-txt-linux-arm64 /usr/local/bin/textWeb

# Copy Template Files
sudo mkdir /var/lib/textweb/templates
sudo mv *.html /var/lib/textweb/templates/

Let textWeb User Own All of Those
sudo chown -R textweb:textweb /var/lib/textweb/*

# Create the SystemD Unit
sudo vim /etc/systemd/system/textweb.service

# Create the Nginx Config
sudo vim /etc/nginx/sites-available/textweb

# Make That Nginx Config Enabled
sudo ln -s /etc/nginx/sites-available/textweb /etc/nginx/sites-enabled/

# Start the textWeb Service
sudo systemctl start textweb.service

# Check that Status
sudo systemctl status textweb.service

# Restart Nginx to Bring in New Config
sudo systemctl restart nginx
```