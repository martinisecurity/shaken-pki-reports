# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate 17599bafbe0c44b8ec89c36eea809196b551d725
Tested At: 2022-10-26 20:19:48 +0000 UTC\
Initial Validity Period: 31 day(s)\
Remaining Validity Period: 10 day(s)\
Subject: CN=SHAKEN Magna5\\, LLC 3849, OU=NOC, O=Magna5\\, LLC, ST=Texas, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/Magna5_3849.crt

View: [Click to view](https://understandingwebpki.com/?cert=MIIDdTCCAxugAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTOMwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAwNTE1MzIwM1oXDTIyMTEwNTE1MzIwM1owYzELMAkGA1UEBhMCVVMxDjAMBgNVBAgMBVRleGFzMRQwEgYDVQQKDAtNYWduYTUsIExMQzEMMAoGA1UECwwDTk9DMSAwHgYDVQQDDBdTSEFLRU4gTWFnbmE1LCBMTEMgMzg0OTBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHSmEvk0ocbvzVCoSD1ktVCycyAFmG%2BqFYMhuhqOwFgUfujidBw3leL%2Fu8Np6stCzm3%2BzRIpi9AhU24BgtokQuGjggGIMIIBhDAWBggrBgEFBQcBGgQKMAigBhYEMzg0OTAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFNQ1JL%2F6yQTbEEZT2L%2BuS97RFERAMIHKBgNVHSMEgcIwgb%2BAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGoYGQpIGNMIGKMQswCQYDVQQGEwJVUzETMBEGA1UECAwKQ2FsaWZvcm5pYTESMBAGA1UEBwwJU2FuIERpZWdvMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEhMB8GA1UEAwwYU0hBS0VOIFNhbnNheSBSb290IENBIFVTghQUtV84BfXwexO1%2BLDe7SDyUXA%2BnDBHBgNVHR8EQDA%2BMDygOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmwwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwCgYIKoZIzj0EAwIDSAAwRQIhAMqr7xOQaQKZ7JHK%2BeWu2I7TZWVGHQo8ORh6%2FoncpBoCAiABHcBA2i3lqDhMl8lqeLQi%2FD4XirUsazPeuj9R05MmoA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 3849' |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |

Generated: 26/10/2022 at 20:21:30