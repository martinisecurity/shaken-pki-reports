# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate c16c41de71faf34e0617e45b50b41db2fee2a460
Tested At: 2022-10-27 22:32:39 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 8 day(s)\
Subject: emailAddress=voiceops@clearfly.net, CN=SHAKEN Greenfly Networks Inc dba Clearfly Communications 210J, OU=NOC, O=Greenfly Networks Inc dba Clearfly Communications, ST=Montana, C=US, emailAddress=voiceops@clearfly.net\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/210J/order/183_210J_63

View: [Click to view](https://understandingwebpki.com/?cert=MIID6TCCA5CgAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTI4wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTAzMjcwM1oXDTIyMTEwNDAzMjcwM1owgdcxCzAJBgNVBAYTAlVTMRAwDgYDVQQIDAdNb250YW5hMTowOAYDVQQKDDFHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zMQwwCgYDVQQLDANOT0MxRjBEBgNVBAMMPVNIQUtFTiBHcmVlbmZseSBOZXR3b3JrcyBJbmMgZGJhIENsZWFyZmx5IENvbW11bmljYXRpb25zIDIxMEoxJDAiBgkqhkiG9w0BCQEWFXZvaWNlb3BzQGNsZWFyZmx5Lm5ldDBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABLd80UROkyOsAvdgaJlDDkYwWMba9%2FmHqLhUC6XioSi0IgY4SJ2NNx9lZWamYeNejzTkK%2FK%2Fhd6a7KVZ2krA5eKjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMjEwSjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFM%2Fv8p05Mv4frtxuMuA9oj7V2coWMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDRwAwRAIgV7%2FwmTXXvHFodKmOWaZN9vkEjXQBOzqVaedWyBYgh2ICIECZKFftOx%2BM8t1gIIuVnTIBYu6O5WlA7BX6YVEzdnY1)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | United States SHAKEN CP | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_ambiguous_identifier | error | United States SHAKEN CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| w_cp_1_3_subject_email | warn | United States SHAKEN CP | Email addresses are not allowed as the CP does not specify how to validate them |
| e_cp1_3_subject_sn | error | United States SHAKEN CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_sti_subject_cn | error | ATIS-1000080 | Common name shall contain the text string 'SHAKEN 210J' |

Generated: 27/10/2022 at 22:33:03