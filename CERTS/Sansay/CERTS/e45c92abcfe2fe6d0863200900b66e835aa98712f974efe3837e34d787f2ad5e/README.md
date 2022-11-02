# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Mitel Cloud Services, Inc. 670J

Tested At: 02 Nov 22 15:09 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 360 day(s)\
Subject: CN=SHAKEN Mitel Cloud Services\\, Inc. 670J, OU=ComNet, O=Mitel Cloud Services\\, Inc., ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/Mitel_Cloud_Services_Inc_670J

View: [Click to view](https://understandingwebpki.com/?cert=MIIDmDCCAz%2BgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkU7IwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyNzIwMTE1NVoXDTIzMTAyNzIwMTE1NVowgYYxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdBcml6b25hMSMwIQYDVQQKDBpNaXRlbCBDbG91ZCBTZXJ2aWNlcywgSW5jLjEPMA0GA1UECwwGQ29tTmV0MS8wLQYDVQQDDCZTSEFLRU4gTWl0ZWwgQ2xvdWQgU2VydmljZXMsIEluYy4gNjcwSjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABA0q7tFRXvcFnnAjZdoOwcKszXHRzLj4IaGVP7WfRjQp57mRQKv%2BH0Ej8OXJazvc6uksjt8cC5MqnSBgO%2Ftyck%2BjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYENjcwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFDQXqBPmwAVEFWsIBckuZWbf%2Fuk8MIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgbUQIeNbUalWtQLFy4pqqNToIU3RxJ9AHtG6TNuBaauICIG2GEZLmZuQJosRCZ1HOilOR0ayZIwh28RGjin9Jytkb)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_atis_subject_cn](../../ISSUES/e_atis_subject_cn/README.md) | error | ATIS1000080 | Common name shall contain the text string 'SHAKEN 670J' |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 02 Nov 22 15:10 UTC