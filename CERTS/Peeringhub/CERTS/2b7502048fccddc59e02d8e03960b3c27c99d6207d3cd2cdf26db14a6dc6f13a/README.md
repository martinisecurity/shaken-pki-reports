# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Celerity Telecom SHAKEN 469K

Tested At: 06 Jul 23 13:57 UTC\
Initial Validity Period: 28 day(s)\
Remaining Validity Period: 27 day(s)\
Subject: CN=Celerity Telecom SHAKEN 469K, O=Celerity Telecom, L=SouthWest Ranches, ST=FL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://certificates.peeringhub.io/469K/469K.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDMTCCAtegAwIBAgIRALI7axf1CLC2IXdfh6bJppIwCgYIKoZIzj0EAwIwfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwHhcNMjMwNzA1MDkyNzQzWhcNMjMwODAxMTU1OTEzWjB4MQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwxGjAYBgNVBAcMEVNvdXRoV2VzdCBSYW5jaGVzMRkwFwYDVQQKDBBDZWxlcml0eSBUZWxlY29tMSUwIwYDVQQDDBxDZWxlcml0eSBUZWxlY29tIFNIQUtFTiA0NjlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAE6XgSP354w8FmDPnDBNkL9Pth%2BmpXJ5nUyXgskpRWXW%2F88MC3SfNqVc1s2BIFoVB3O3U0SGMM6pXUw0qmMtx9WaOCATwwggE4MA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBT50vn46bILMh%2Fk7TvH3PJVIAsnqDAfBgNVHSMEGDAWgBSuoXNRiClXEcoMqfSxCm5OuEtNBzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwFgYIKwYBBQUHARoECjAIoAYWBDQ2OUswgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0gAMEUCIAF1P5lwRBmbFuqdK5Ia%2Bh%2FCs%2FeeHIopvlLZJu0IKrpDAiEAkW8rU62MbhDrrNnNCnUfyVcjDrcDw2FPu04FYAO5pQA%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |


Generated: 06 Jul 23 14:08 UTC