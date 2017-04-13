# Boshler

Like bundler for bosh releases...kinda.

Boshler caches releases defined in a BOSHFILE locally and uploads them to the currently targeted BOSH director.
Releases are cached at $HOME/.boshler/releases and are downloaded from http://bosh.io as needed.  

## Install

```
go install github.com/zrob/boshler
```

## Usage

```
boshler
```

## Example BOSHFILE

```json
{
  "releases": [
    {
      "name": "garden-runc-release",
      "repository": "cloudfoundry"
    },
    {
      "name": "etcd-release",
      "repository": "cloudfoundry-incubator"
    },
    {
      "name": "cflinuxfs2-rootfs-release",
      "repository": "cloudfoundry"
    }
  ]
}
```
