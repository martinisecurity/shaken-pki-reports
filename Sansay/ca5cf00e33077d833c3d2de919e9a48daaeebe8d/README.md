# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate ca5cf00e33077d833c3d2de919e9a48daaeebe8d
Tested At: 2022-10-26 21:00:46 +0000 UTC\
Initial Validity Period: 365 day(s)\
Remaining Validity Period: 350 day(s)\
Subject: CN=SHAKEN ALD Telecom 780J, OU=ALD_Telecom, O=ALD Telecom, ST=Arizona, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/ALD_Telecom_780J

View: [Click to view](https://understandingwebpki.com/?cert=MIIC0TCCAnegAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTwkwCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MDg0M1oXDTIzMTAxMTE3MDg0M1owbTELMAkGA1UEBhMCVVMxEDAOBgNVBAgMB0FyaXpvbmExFDASBgNVBAoMC0FMRCBUZWxlY29tMRQwEgYDVQQLDAtBTERfVGVsZWNvbTEgMB4GA1UEAwwXU0hBS0VOIEFMRCBUZWxlY29tIDc4MEowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAATrCf7p2dfITkSwC6B1rTvmhygSeQJPCk%2F%2BrBOtuvbh5Xq4ZxNiagKMfjIcD3ttnBAUoy%2BprtYjRyGMsg0zaLXIo4HbMIHYMBYGCCsGAQUFBwEaBAowCKAGFgQ3ODBKMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBATAdBgNVHQ4EFgQU5D6yJDqBy8u7UBJwCsnFyuynU18wHwYDVR0jBBgwFoAUrNOT9UNDzAq%2BRVgXE32SfNzDAUYwRwYDVR0fBEAwPjA8oDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsMAwGA1UdEwEB%2FwQCMAAwDgYDVR0PAQH%2FBAQDAgeAMAoGCCqGSM49BAMCA0gAMEUCIGCAha%2FoKlqwoEUEXitI977nsKa1Sp3XOapSLkKeNzOUAiEAzSRGlMcIq%2By8WKxK8YEeYpeuJBp5IS7gpRJCsdkQDys%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 780J' |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |

Generated: 26/10/2022 at 21:01:13