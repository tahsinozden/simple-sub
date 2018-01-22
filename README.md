# tubtitle
a simple subtitle tool which can remove accent characters, merge subtitles etc. The tool is made up of two parts. First part is a command-line interface. Second one is a web application. Both interfaces provide the same functionalities.

## how-to-build
Run the `build.sh` script to build the apllication. `tubtitle` binary will be generated. `./tubtitle -help` can provide the all options inside the application.

## how-to-remove-accent-characters
`./tubtitle -mode remove-accent -f sub.srt -enc pl` 

## how-to-merge-subtitles
`./tubtitle -mode merge -f1 sub1.srt -e1 pl -f2 sub2.srt -e2 tr` 

## how-to-run-webapp
`./tubtitle -mode serve` the application will run on the default port `3000`. i.e. `http://localhost:3000` <br>
`./tubtitle -mode serve -p <your_port>` the application will run in a specified port. <br>
below image is a sample from the application web interface. <br>
![img](https://github.com/tahsinozden/tubtitle/blob/master/static/main_page.png)
