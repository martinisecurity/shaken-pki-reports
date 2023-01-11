# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ytel Inc. 703J

Tested At: 11 Jan 23 21:52 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 281 day(s)\
Subject: CN=SHAKEN Ytel Inc. 703J, OU=Information Technology, O=Ytel Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://s3.amazonaws.com/certificates.peeringhub.io/ytel/703J.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDiTCCAy%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUd0wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxOTE5MDg1OVoXDTIzMTAxOTE5MDg1OVowdzELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAoMCVl0ZWwgSW5jLjEfMB0GA1UECwwWSW5mb3JtYXRpb24gVGVjaG5vbG9neTEeMBwGA1UEAwwVU0hBS0VOIFl0ZWwgSW5jLiA3MDNKMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEUq7k4svt6ztcXyH%2FDdEaJYj7XVfAWgWRfiiRhM%2BrYP5AU7K4%2FPCvhh2H5pEGOLvuoIl%2F3pgi8TelAG%2Bi7l%2BB96OCAYgwggGEMBYGCCsGAQUFBwEaBAowCKAGFgQ3MDNKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQUS0Z5E29XusSq0D4hGtTr%2FI52M8cwgcoGA1UdIwSBwjCBv4AUrNOT9UNDzAq%2BRVgXE32SfNzDAUahgZCkgY0wgYoxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRIwEAYDVQQHDAlTYW4gRGllZ28xGzAZBgNVBAoMElNhbnNheSBDb3Jwb3JhdGlvbjESMBAGA1UECwwJU2Fuc2F5IENBMSEwHwYDVQQDDBhTSEFLRU4gU2Fuc2F5IFJvb3QgQ0EgVVOCFBS1XzgF9fB7E7X4sN7tIPJRcD6cMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNIADBFAiBdH4lt9vOP1B0Or2CyO1TQvO6XKiQ4xRw1PSELtdPXewIhALT3Hgn1%2BdV5uTotUu3sBe%2Fu8eaHMoMN%2FkQiuhEa9HYO)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 703J' |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 11 Jan 23 21:59 UTC