# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN ConnectMeVoice 719J

Tested At: 31 Oct 22 16:41 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 348 day(s)\
Subject: CN=SHAKEN ConnectMeVoice 719J, OU=Production, O=ConnectMe\\, LLC, ST=New Jersey, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/ConnectMeVoice

View: [Click to view](https://understandingwebpki.com/?cert=MIIC2jCCAn%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUAUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMzIwMjc1N1oXDTIzMTAxMzIwMjc1N1owdTELMAkGA1UEBhMCVVMxEzARBgNVBAgMCk5ldyBKZXJzZXkxFzAVBgNVBAoMDkNvbm5lY3RNZSwgTExDMRMwEQYDVQQLDApQcm9kdWN0aW9uMSMwIQYDVQQDDBpTSEFLRU4gQ29ubmVjdE1lVm9pY2UgNzE5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABP7NBXnnL4PImD%2F5VtqomjjL3HYEJHJXfB5kNI0Y26ztyUdv1k08cysGRlsawwEg1b%2Fv%2FldddKtjQSgTHnJ%2BzG2jgdswgdgwFgYIKwYBBQUHARoECjAIoAYWBDcxOUowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEBMB0GA1UdDgQWBBSFkm9ICV8rZ37Q0gIySHRXPtk5dzAfBgNVHSMEGDAWgBSs05P1Q0PMCr5FWBcTfZJ83MMBRjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhALBrOwD6bgnggoxIBRMwCD3FScJud5MH02bI6G7Ee%2BjaAiEA%2BB1%2BeHQHzjm65gw93Z%2F4uG%2FIZzqTmUnGr1HeNJngPbI%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 719J' |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 31/10/2022 at 16:43:22