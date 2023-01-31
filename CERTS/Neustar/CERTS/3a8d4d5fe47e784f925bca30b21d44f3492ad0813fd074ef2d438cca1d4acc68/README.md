# STIR/SHAKEN CA Ecosystem Compliance

## Certificate ATT SHAKEN 4036

Tested At: 31 Jan 23 17:08 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 211 day(s)\
Subject: CN=ATT SHAKEN 4036, OU=AT&T Services\\, Inc., O=ATT SHAKEN E-E, L=Dallas, ST=Texas, C=US\
Issuer: CN=Neustar Certified Caller ID SHAKEN CA-1, OU=www.ccid.neustar, O=Neustar Information Services Inc, C=US\
Link: https://cert.sticr.att.net:8443/certs/att/abbf5398-e1e1-42af-96a7-092303b168ba

[View certificate details](https://understandingwebpki.com/?cert=MIIDQjCCAuegAwIBAgIUFriSAbdqxX6d3o0zKLwnhEOum00wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMSkwJwYDVQQKDCBOZXVzdGFyIEluZm9ybWF0aW9uIFNlcnZpY2VzIEluYzEZMBcGA1UECwwQd3d3LmNjaWQubmV1c3RhcjEwMC4GA1UEAwwnTmV1c3RhciBDZXJ0aWZpZWQgQ2FsbGVyIElEIFNIQUtFTiBDQS0xMB4XDTIyMDgyOTE5MDcxNFoXDTIzMDgyOTE5MDcxNFowfzELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMQ8wDQYDVQQHDAZEYWxsYXMxFzAVBgNVBAoMDkFUVCBTSEFLRU4gRS1FMRwwGgYDVQQLDBNBVCZUIFNlcnZpY2VzLCBJbmMuMRgwFgYDVQQDDA9BVFQgU0hBS0VOIDQwMzYwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARDggzP38GpGdPr4F0YMI%2FNSGC3rXY7ti4n334uc1u9NEgtFp%2Fiot4TPT3fw27QhKwuJ1ZV%2FLu3IzIOfkJK6b68o4IBODCCATQwFgYIKwYBBQUHARoECjAIoAYWBDQwMzYwDAYDVR0TAQH%2FBAIwADAfBgNVHSMEGDAWgBSv0cjC7nJMg%2Fw%2F7RmnbR2QsgfwOjBaBggrBgEFBQcBAQROMEwwSgYIKwYBBQUHMAKGPmh0dHA6Ly9jYWNlcnRzLXVzLmNjaWQubmV1c3Rhci9OZXVzdGFyQ2VydGlmaWVkQ2FsbGVySWRDQTEuY3J0MBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwHQYDVR0OBBYEFDw5j7%2B473%2BTcPBwwb6YRh25GFNDMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNJADBGAiEA51j0MSEOOWHU%2B13Nw1OOUpe10A1DasTA%2BlgJCPMLg%2BECIQD5%2BMdL0D%2FhCLTwvU1u8X3QFXxJE0vYvC0ue8nESmcHug%3D%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_atis_extension_unknown](../../ISSUES/e_atis_extension_unknown/README.md) | error | ATIS1000080 | STI certificate shall not include extensions that are not specified |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 31 Jan 23 17:51 UTC