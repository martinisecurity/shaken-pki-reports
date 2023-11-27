# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN DLS Internet Services 815J

Tested At: 27 Nov 23 22:45 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 2 day(s)\
Subject: CN=SHAKEN DLS Internet Services 815J, OU=sansaynss01.dls.net, O=DLS Internet Services, ST=Illinois, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/DLS_Internet_Services_815J

[View certificate details](https://understandingwebpki.com/?cert=MIIC8DCCApWgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkhdQwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMTAzMDE0MDc0M1oXDTIzMTEyOTE0MDc0M1owgYoxCzAJBgNVBAYTAlVTMREwDwYDVQQIDAhJbGxpbm9pczEeMBwGA1UECgwVRExTIEludGVybmV0IFNlcnZpY2VzMRwwGgYDVQQLDBNzYW5zYXluc3MwMS5kbHMubmV0MSowKAYDVQQDDCFTSEFLRU4gRExTIEludGVybmV0IFNlcnZpY2VzIDgxNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQ%2FL5g4yQSabbyTek%2FaTQM3Dzu%2F6oFmgPLKn9y%2FG%2FzDwzO5M9ob4U2nuciqD9IEUephnlJ7Y2jFO5AY%2B69Dtwkso4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ4MTVKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU0wyRK%2BVLlV11Hq0JUj2f2Mu1gIEwHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQDR0RDiP6DRvY4rj1GCbIYDKKSLxgMjZQBo1JQ5tBZtIgIhAO7bnkfLnkaPyEreJMzJdl697VX6s%2B1Lf1fBYMfbQVmv)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 815J', but common name is 'SHAKEN DLS Internet Services 815J' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |


Generated: 27 Nov 23 22:56 UTC