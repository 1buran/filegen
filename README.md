# File generator
[![goreportcard](https://goreportcard.com/badge/github.com/1buran/filegen)](https://goreportcard.com/report/github.com/1buran/filegen)
![Demo](https://i.imgur.com/6SeGc15.gif)

> [!CAUTION]
> Script is optimized to generate a lot of small text files with random chars within,
> please do not use it for generation files with size more than tens of megabytes,
> it will not be effective cos it keep in memory whole content of file.
> For generation of big files please consider usage of `/dev/random` or `/dev/zero`,
> for example: `dd if=/dev/random of=1Gb-file bs=1G count=1`

# Usage

Installation:
```
$ go install github.com/1buran/filegen@latest
```

Create 1 hundred of 10K files in `test1` dir:
```
$ filegen -count 100 -size 10K -path test1
```

Create 1 thousand of files with random size from 5K to 10K:
```
$ filegen -count 1000 -size 10K -random-size-min 5K
```

## Tasks

These are tasks of [xc](https://github.com/joerdav/xc) runner.

### vhs

Run VHS fo update gifs.

```
vhs demo.tape
```

### imgur

Upload to Imgur and update readme.
```
url=`curl --location https://api.imgur.com/3/image \
     --header "Authorization: Client-ID ${clientId}" \
     --form image=@demo.gif \
     --form type=image \
     --form title=filegen \
     --form description=Demo | jq -r '.data.link'`
sed -i "s#^\!\[Demo\].*#![Demo]($url)#" README.md
```
