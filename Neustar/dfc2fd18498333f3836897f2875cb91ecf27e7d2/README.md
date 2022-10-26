# STIR/SHAKEN CA Ecosystem Compliance
## Neustar

### Certificate dfc2fd18498333f3836897f2875cb91ecf27e7d2
Tested At: 2022-10-26 23:14:33 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 167 day(s)\
Subject: CN=SHAKEN 951J, O=Zultys Inc, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US

Link: https://zultys-pem-cert-2022.s3.amazonaws.com/89179fa533bbaf0aea20a9f31aa06f20.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC%2BzCCAqGgAwIBAgIUDraMtHqPxPkd5RlOE6DklUQcVyswCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDQxMTE4MzM0OFoXDTIzMDQxMTE4MzM0OFowODELMAkGA1UEBhMCVVMxEzARBgNVBAoMClp1bHR5cyBJbmMxFDASBgNVBAMMC1NIQUtFTiA5NTFKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEsQk3BYPAeFI9cbM18%2BBWNM8roF2vH1zl2jmB0AlJiiOUDK29PgSBLgLDRa8e6y0UB5KBmUrxJRBSvqPsUpMmBqOCATkwggE1MBYGCCsGAQUFBwEaBAowCKAGFgQ5NTFKMAwGA1UdEwEB%2FwQCMAAwHwYDVR0jBBgwFoAUr9HIwu5yTIP8P%2B0Zp20dkLIH8DowWwYIKwYBBQUHAQEETzBNMEsGCCsGAQUFBzAChj9odHRwOi8vY2FjZXJ0cy11cy5jY2lkLm5ldXN0YXIvTmV1c3RhckNlcnRpZmllZENhbGxlcklkQ0ExLmNydCAwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAdBgNVHQ4EFgQUINiLYo980pp3LIz%2FObth9jJ0ggswDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIQDSFGTkEiJkLWPlHgU1yMEjWw8tjdEUjO3YiYr8gMxwsgIgah2n%2B01Ci7S3kfvZwbTiq1cEMxNGul1n7W7ubqd0qkA%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_extension_unknown | error | ATIS-1000080v4 | STI certificate shall not include extensions that are not specified |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |

Generated: 26/10/2022 at 23:14:41