# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Lightspeed Voice 557F

Tested At: 21 Nov 23 01:42 UTC\
Initial Validity Period: 60 day(s)\
Remaining Validity Period: -401 day(s)\
Subject: emailAddress=level5@lightspeedvoice.com, CN=SHAKEN Lightspeed Voice 557F, OU=NOC, O=Lightspeed Voice, ST=Florida, C=US, emailAddress=level5@lightspeedvoice.com\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/557F/order/63_557F_83

[View certificate details](https://understandingwebpki.com/?cert=MIIDrjCCA1OgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkQeEwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDgxNjExNDM0MVoXDTIyMTAxNTExNDM0MVowgZoxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdGbG9yaWRhMRkwFwYDVQQKDBBMaWdodHNwZWVkIFZvaWNlMQwwCgYDVQQLDANOT0MxJTAjBgNVBAMMHFNIQUtFTiBMaWdodHNwZWVkIFZvaWNlIDU1N0YxKTAnBgkqhkiG9w0BCQEWGmxldmVsNUBsaWdodHNwZWVkdm9pY2UuY29tMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBXl3GCua%2BqMXp31Y8k4f0nN4gH7VMfG4sDv%2FiicinVgnyuVMVKGxE6ogbymRuI31eSsPawjrSItpvtyAgyUGcKOCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ1NTdGMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUVyKFWLDI74e8KWf8JuXGw9Db0f0wgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6aMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEAuzSIShvA6htC%2FW%2FpwX42Tc4yT%2BDC8U%2B4Hi%2BLes8Mx0UCIQDWo6xf1HX66BSM3Uv9AAeu%2BzjLirk%2FOkCpCruAF0SLSw%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_shaken_certificate_policies_id](../../ISSUES/e_shaken_certificate_policies_id/README.md) | error | US_SHAKEN_CP | the Certificate Policies extension contains an invalid OID value: 2.16.840.1.114569.1.1.1. Available OIDs: [2.16.840.1.114569.1.1.3 2.16.840.1.114569.1.1.4] |
| [e_atis_subject_cn_spc](../../ISSUES/e_atis_subject_cn_spc/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 557F' |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC