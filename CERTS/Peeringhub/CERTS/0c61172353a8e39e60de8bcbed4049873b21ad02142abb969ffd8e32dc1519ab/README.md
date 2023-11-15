# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Celerity Telecom SHAKEN 469K

Tested At: 15 Nov 23 18:05 UTC\
Initial Validity Period: 20 day(s)\
Remaining Validity Period: -6 day(s)\
Subject: CN=Celerity Telecom SHAKEN 469K, O=Celerity Telecom, L=SouthWest Ranches, ST=FL, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://s3.amazonaws.com/certificates.peeringhub.io/469K/469K.crt

[View certificate details](https://understandingwebpki.com/?cert=MIIDMTCCAtegAwIBAgIRALPF4h3O5C21zUxsViWhZTUwCgYIKoZIzj0EAwIwfDELMAkGA1UEBhMCVVMxFzAVBgNVBAoMDlBlZXJpbmdodWIgSW5jMSIwIAYDVQQLDBlDZXJ0aWZpY2F0aW9uIEF1dGhvcml0aWVzMTAwLgYDVQQDDCdQZWVyaW5naHViIEluYyBTSEFLRU4gSW50ZXJtZWRpYXRlIENBIDIwHhcNMjMxMDIwMTkzNTEyWhcNMjMxMTA5MTUzOTM0WjB4MQswCQYDVQQGEwJVUzELMAkGA1UECAwCRkwxGjAYBgNVBAcMEVNvdXRoV2VzdCBSYW5jaGVzMRkwFwYDVQQKDBBDZWxlcml0eSBUZWxlY29tMSUwIwYDVQQDDBxDZWxlcml0eSBUZWxlY29tIFNIQUtFTiA0NjlLMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAErK1ulFRz%2FjDNULkXmMW5l0M5H8rxWT2mZemH9iMYL87QUGg9MDo3sGFvM80hQ0rqUa3svhy9ZeAkQhyD4UdSPaOCATwwggE4MA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBQjHTm%2F8P2fzIDIzhtHd8YcKFVeQjAfBgNVHSMEGDAWgBSuoXNRiClXEcoMqfSxCm5OuEtNBzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwFgYIKwYBBQUHARoECjAIoAYWBDQ2OUswgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0gAMEUCIC8rWNzMNj4kytW0R7nW3crBE9hZ4uJo5OJPkma8CvT8AiEAn%2F0YCg%2F%2FCTSqnMqVTEOeh8sLmc3X3EwXkRf6CtYW1AQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |


Generated: 15 Nov 23 18:10 UTC