# STIR/SHAKEN CA Ecosystem Compliance
## Comcast

### Certificate f9fdf7a4cb279213b0cb0d2f0638caa19dd4c36c
Tested At: 2022-10-26 23:14:39 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 24 day(s)\
Subject: CN=SHAKEN, O=Comcast, L=Philadelphia, ST=Pennsylvania, C=US\
Issuer: CN=Comcast SHAKEN Intermediate CA, O=Comcast, ST=Pennsylvania, C=US

Link: https://sticr.stir.comcast.com/7fa7ab43-191d-4a8d-ac17-ebed13b7b30d.cer

View: [Click to view](https://understandingwebpki.com/?cert=MIICVjCCAf2gAwIBAgIITWWCIQf8%2FVIwCgYIKoZIzj0EAwIwXzELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEQMA4GA1UECgwHQ29tY2FzdDEnMCUGA1UEAwweQ29tY2FzdCBTSEFLRU4gSW50ZXJtZWRpYXRlIENBMB4XDTIyMTAyMDExMTIyM1oXDTIyMTExOTExMTIyM1owXjELMAkGA1UEBhMCVVMxFTATBgNVBAgTDFBlbm5zeWx2YW5pYTEVMBMGA1UEBxMMUGhpbGFkZWxwaGlhMRAwDgYDVQQKEwdDb21jYXN0MQ8wDQYDVQQDEwZTSEFLRU4wWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAR3EawXmkkgYnHqMjZorkLIGnXhG0NnGcXMOHUzzqyTWP819jX%2BR%2FmfiWxeuKAM6WD%2BJ6ft7O7lAlWvQ7MWII8Ao4GjMIGgMA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB8GA1UdIwQYMBaAFL%2BOHGBVJMURSM9CRUS%2FsPQ6K%2BYoMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAWBggrBgEFBQcBGgQKMAigBhYEMzE4SjAKBggqhkjOPQQDAgNHADBEAiBit4eayNM%2BM%2FOC1ceEntSIPyqdnjb0XQnIY0kf35N6kAIgFujhldfYNGdsEKufcoxemL%2BHf4Mq3DoSasS14BEcUxU%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 318J' |
| w_ext_subject_key_identifier_missing_sub_cert | warn | RFC5280 |  |
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_subject_key_identifier | error | ATIS-1000080v4 | STI certificates shall contain a Subject Key Identifier extension |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall include a Certificate Policies extension containing a single SHAKEN Certificate Policy |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 23:14:41