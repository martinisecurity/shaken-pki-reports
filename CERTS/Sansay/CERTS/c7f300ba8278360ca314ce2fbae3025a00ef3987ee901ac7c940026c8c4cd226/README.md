# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J

Tested At: 31 Oct 22 18:33 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 20 day(s)\
Subject: emailAddress=voiceops@clearfly.net, CN=SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J, OU=NOC, O=Greenfly Networks Inc dba Clearfly Communications, ST=Montana, C=US, emailAddress=voiceops@clearfly.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US\
Link: https://cr.sansay.com/210J/order/199_210J_63

View: [Click to view](https://understandingwebpki.com/?cert=MIID6jCCA5CgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkUhgwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAyMTAyMDcxNloXDTIyMTEyMDAyMDcxNlowgdcxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdNb250YW5hMTowOAYDVQQKDDFHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zMQwwCgYDVQQLDANOT0MxRjBEBgNVBAMMPVNIQUtFTiBHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zIDIxMEoxJDAiBgkqhkiG9w0BCQEWFXZvaWNlb3BzQGNsZWFyZmx5Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLd80UROkyOsAvdgaJlDDkYwWMba9%2FmHqLhUC6XioSi0IgY4SJ2NNx9lZWamYeNejzTkK%2FK%2Fhd6a7KVZ2krA5eKjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMjEwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFM%2Fv8p05Mv4frtxuMuA9oj7V2coWMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAPgO1wxKfNyWqEwtPfmLc7w40pe6JXgIFkkQRF2Mm2dyAiBf4KcuLuf6Au5it%2B3HVW4xRna6iPVADyz53K1Q9b%2B1uQ%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_pki_subject_rdn_unknown](../../ISSUES/w_pki_subject_rdn_unknown/README.md) | warn | SHAKEN PKI Best Practice | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_sti_subject_cn](../../ISSUES/e_sti_subject_cn/README.md) | error | ATIS&#x2011;1000080 | Common name shall contain the text string 'SHAKEN 210J' |
| [e_cp1_3_ambiguous_identifier](../../ISSUES/e_cp1_3_ambiguous_identifier/README.md) | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_cp1_3_subject_sn](../../ISSUES/e_cp1_3_subject_sn/README.md) | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [w_cp_1_3_subject_email](../../ISSUES/w_cp_1_3_subject_email/README.md) | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| [e_sti_certificate_policies](../../ISSUES/e_sti_certificate_policies/README.md) | error | ATIS&#x2011;1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 31/10/2022 at 18:34:12