# STIR/SHAKEN CA Ecosystem Compliance
## Sansay

### Certificate d6e72b06e603a6bf28e159db14e770a48c86cc58
Tested At: 2022-10-26 21:00:53 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 15 day(s)\
Subject: CN=SHAKEN VoIP Innovations 597F, OU=NOC, O=VoIP Innovations, ST=Pennsylvania, C=US\
Issuer: CN=SHAKEN Sansay Intermediate CA US WEST 1, OU=Sansay CA, O=Sansay Corporation, ST=California, C=US

Link: https://cr.sansay.com/VoIP_Innovations_597F

View: [Click to view](https://understandingwebpki.com/?cert=MIIC1zCCAn6gAwIBAgIUQpx8cHEeOCDwuOHerm%2FzJiJkTz8wCgYIKoZIzj0EAwIwgYUxCzAJBgNVBAYTAlVTMRMwEQYDVQQIDApDYWxpZm9ybmlhMRswGQYDVQQKDBJTYW5zYXkgQ29ycG9yYXRpb24xEjAQBgNVBAsMCVNhbnNheSBDQTEwMC4GA1UEAwwnU0hBS0VOIFNhbnNheSBJbnRlcm1lZGlhdGUgQ0EgVVMgV0VTVCAxMB4XDTIyMTAxMTE3MjMyMloXDTIyMTExMDE3MjMyMlowdDELMAkGA1UEBhMCVVMxFTATBgNVBAgMDFBlbm5zeWx2YW5pYTEZMBcGA1UECgwQVm9JUCBJbm5vdmF0aW9uczEMMAoGA1UECwwDTk9DMSUwIwYDVQQDDBxTSEFLRU4gVm9JUCBJbm5vdmF0aW9ucyA1OTdGMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEZxb5rNCMqQTKIowEPxI78IPaRDMoC1gJoQtw%2FiJ%2BjTAuI2%2F4nDDGBz7ACAo3eIerjz3WANlpeFOQbeXW03U8jaOB2zCB2DAWBggrBgEFBQcBGgQKMAigBhYENTk3RjAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwHQYDVR0OBBYEFOKfbEv6IldN3qqnsd3NpPk2C9PwMB8GA1UdIwQYMBaAFKzTk%2FVDQ8wKvkVYFxN9knzcwwFGMEcGA1UdHwRAMD4wPKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybDAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAKBggqhkjOPQQDAgNHADBEAiAL%2FVD7QaqL0gjyPDlsHh0lhwiuyNjPyhw7Y7bMB%2BJWpgIgXMh%2FEZrVvjD9aQuo%2B54tZXUTZGWtw7UsukHRZmRSid4%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| e_cp1_3_ambiguous_identifier | error | CPv1.3 | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_sti_certificate_policies | error | ATIS-1000080v4 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_subject_cn | error | ATIS-1000080v4 | Common name shall contain the text string 'SHAKEN 597F' |
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |

Generated: 26/10/2022 at 21:01:13