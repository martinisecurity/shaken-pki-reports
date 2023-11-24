# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Cyberlynk Network, LLC 086K

Tested At: 24 Nov 23 11:07 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: -291 day(s)\
Subject: emailAddress=software@cyberlynk.net, CN=SHAKEN Cyberlynk Network\\, LLC 086K, OU=NOC, O=Cyberlynk Network\\, LLC, ST=Winsconsin, C=US, emailAddress=software@cyberlynk.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/086K/order/203_086K_91

[View certificate details](https://understandingwebpki.com/?cert=MIIDCjCCArCgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkW1kwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIzMDEwNjE0NDI1NloXDTIzMDIwNTE0NDI1NlowgaUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApXaW5zY29uc2luMR8wHQYDVQQKDBZDeWJlcmx5bmsgTmV0d29yaywgTExDMQwwCgYDVQQLDANOT0MxKzApBgNVBAMMIlNIQUtFTiBDeWJlcmx5bmsgTmV0d29yaywgTExDIDA4NksxJTAjBgkqhkiG9w0BCQEWFnNvZnR3YXJlQGN5YmVybHluay5uZXQwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATpkCFA8LvfWu1asIDP6BOFqM1qDdJaPFG22XbgahglidWLpWI5TFVs9ve8O2SFujBOe8D%2FL3v33OCGvktdD0iPo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQwODZLMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQULPPQSRAOl8C1R02jCqCfhikI2SswHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIGpH57ThScRByfXkeszwThquslDX5QREOv7InQ3ECjH1AiEAn0wM7j9ZDaZHuHB1Cu11wGZ7zchz%2FJuZGZlJW5VGJgI%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 086K', but common name is 'SHAKEN Cyberlynk Network, LLC 086K' |
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | The Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: 2.16.840.1.114569.1.1.3, 2.16.840.1.114569.1.1.4 |
| [e_atis_ext_crl_distribution_struct](../../ISSUES/e_atis_ext_crl_distribution_struct/README.md) | error | ATIS1000080 | CRL Distribution Point shall contain a CRLIssuer field |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 24 Nov 23 11:17 UTC