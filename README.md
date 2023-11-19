# go-tso

TiDB cli tool to convert timestamps to human strings

# Usage

The cli tool can accept a timestamp in an almost RFC3339 as well as a TiDB timestamp object and
will output the other.


Convert a timestamp object to RFC3339 format

```bash
$ tso 445544556134400000
  Converting input '445544556134400000' to RFC3339...
  2023-11-10T12:00:00Z

$ tso 2023-11-10T12:00:00Z
  Converting input '2023-11-10T12:00:00Z' to TSO...
  445544556134400000
```

The program can also take the input from stdin if processing data

```bash
$ echo 445555063652352000 | tso
  Converting input '445555063652352000' to RFC3339...
  2023-11-10T23:08:03Z

$ kubectl get backup log-backup -o json | jq .status.logSuccessTruncateUntil | tso
  Converting input '2022-01-17T12:42:05Z' to TSO...
  430551420108800000

$ echo 2023-11-01T12:00:00Z 2023-11-02T12:00:00Z 2023-11-03T12:00:00Z 2023-11-04T12:00:00Z | xargs -n 1 tso | tee output.txt
  Converting input '2023-11-01T12:00:00Z' to TSO...
  445340712960000000
  Converting input '2023-11-02T12:00:00Z' to TSO...
  445363362201600000
  Converting input '2023-11-03T12:00:00Z' to TSO...
  445386011443200000
  Converting input '2023-11-04T12:00:00Z' to TSO...
  445408660684800000

$ cat output.txt
  445340712960000000
  445363362201600000
  445386011443200000
  445408660684800000
```
