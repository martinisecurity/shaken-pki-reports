# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN InteractiveTel, LLC 920J

Tested At: 27 Nov 23 23:20 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -260 day(s)\
Subject: emailAddress=jfindley@interactivetel.com, CN=SHAKEN InteractiveTel\\, LLC 920J, OU=NOC, O=InteractiveTel\\, LLC, ST=Texas, C=US, emailAddress=jfindley@interactivetel.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/920J/order/238_920J_66

[View certificate details](https://understandingwebpki.com/?cert=MIIDBDCCAqqgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkYRcwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDIxMDE0NDcwMloXDTIzMDMxMjE0NDcwMlowgZ8xCzAJBgNVBAYTAlVTMQ4wDAYDVQQIDAVUZXhhczEcMBoGA1UECgwTSW50ZXJhY3RpdmVUZWwsIExMQzEMMAoGA1UECwwDTk9DMSgwJgYDVQQDDB9TSEFLRU4gSW50ZXJhY3RpdmVUZWwsIExMQyA5MjBKMSowKAYJKoZIhvcNAQkBFhtqZmluZGxleUBpbnRlcmFjdGl2ZXRlbC5jb20wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASMcpDf9PfrQJvZ2I%2F245uHQmrHEDVHkPOeja8uEeyKGGQmCO5Z%2BIyec1ynTksWHveCDCq4BKha7Ft4sRx5xmLUo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ5MjBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzAdBgNVHQ4EFgQUTOmRudIlU1mpp%2FGRM%2B2axh62DtowHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCID2z75iF0RvhA2pBqyFdKkm5rDQwWitZ21jGiRtTbJ4LAiEAh%2BstM2DGviss3wDXHIF46nzpwkvSpr18Zlov%2FgHr0%2BA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 920J', but common name is 'SHAKEN InteractiveTel, LLC 920J' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 27 Nov 23 23:28 UTC