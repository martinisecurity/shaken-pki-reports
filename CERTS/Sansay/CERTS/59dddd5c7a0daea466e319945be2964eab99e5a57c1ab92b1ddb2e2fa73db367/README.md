# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Ytel Inc. 703J

Tested At: 02 Dec 22 07:45 UTC\
Initial Validity Period: 130 day(s)\
Remaining Validity Period: -51 day(s)\
Subject: CN=SHAKEN Ytel Inc. 703J, OU=voice, O=Ytel Inc., ST=California, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://s3.amazonaws.com/certificates.peeringhub.io/ytel/703J.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDeTCCAx6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkOTUwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMDYwMzIyMjQxOVoXDTIyMTAxMTIyMjQxOVowZjELMAkGA1UEBhMCVVMxEzARBgNVBAgMCkNhbGlmb3JuaWExEjAQBgNVBAoMCVl0ZWwgSW5jLjEOMAwGA1UECwwFdm9pY2UxHjAcBgNVBAMMFVNIQUtFTiBZdGVsIEluYy4gNzAzSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKco4XZU966SB1Rj4unGediDr9su3Yb95bO5%2FzzukzboadgxbhsP0mGKbTjGmG8sPGPK7nZRqat3B2u6ydfxDXCjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENzAzSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDPhXure7G2QobiIsc4JnXnK%2FH%2FuMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BmjBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSQAwRgIhALru7LYo0oRUydA2kTaxoOA%2Ffn6lHkyUHNXMknoDDNj4AiEAsT7vzF4IUIDnQRlHDbohuI2Euk85n8qRWsyakPC3TWM%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 703J' |


Generated: 02 Dec 22 07:46 UTC