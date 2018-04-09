import time
import subprocess

minlat = 54.6647500 
minlon = 55.7748500 
maxlat = 54.8314500 
maxlon = 56.1668500

numberOfIteration = 0
allTime = 0

lat = minlat
while lat < maxlat:
    lon = minlon
    while lon < maxlon:
        lat = round(lat, 7)
        lon = round(lon, 7)

        file = open('input/Task2/point.xml', 'w')
        file.write('<?xml version="1.0" encoding="UTF-8"?>\n')
        file.write('<point lat="' + str(lat) + '" lon="' + str(lon) + '" />\n')
        file.close()

        startTime = time.time()

        subprocess.call('go run $GOPATH/src/github.com/vahriin/BigGraph/Task2/Task2.go -t', shell=True)

        endTime = time.time()

        allTime += (endTime - startTime)

        lon += (maxlon - minlon) / 10
    lat += (maxlat - minlat) / 10

print(allTime)