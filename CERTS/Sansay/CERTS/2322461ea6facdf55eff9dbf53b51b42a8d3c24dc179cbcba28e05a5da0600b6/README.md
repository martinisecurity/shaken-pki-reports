# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InteractiveTel, LLC 920J

Tested At: 04 Oct 24 16:13 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -653 day(s)\
Subject: emailAddress=jfindley@interactivetel.com, CN=SHAKEN InteractiveTel\\, LLC 920J, OU=NOC, O=InteractiveTel\\, LLC, ST=Texas, C=US, emailAddress=jfindley@interactivetel.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/920J/order/156_920J_66

[View certificate details](https://x509.io/?cert=MIIDBTCCAqqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkVgkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTEyMTExNTA1MFoXDTIyMTIyMTExNTA1MFowgZ8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczEcMBoGA1UECgwTSW50ZXJhY3RpdmVUZWwsIExMQzEMMAoGA1UECwwDTk9DMSgwJgYDVQQDDB9TSEFLRU4gSW50ZXJhY3RpdmVUZWwsIExMQyA5MjBKMSowKAYJKoZIhvcNAQkBFhtqZmluZGxleUBpbnRlcmFjdGl2ZXRlbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASMcpDf9PfrQJvZ2I%2F245uHQmrHEDVHkPOeja8uEeyKGGQmCO5Z%2BIyec1ynTksWHveCDCq4BKha7Ft4sRx5xmLUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUTOmRudIlU1mpp%2FGRM%2B2axh62DtowHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0kAMEYCIQD7xJfYL8iSK5OwkTKraEnYPDoApoALpgW6eFUfOQ%2BDXwIhAP3hWGpkjPan3kzqWonqKM5MSky1x3aqkzx2HEAUWTUb)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 920J', but common name is 'SHAKEN InteractiveTel, LLC 920J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC