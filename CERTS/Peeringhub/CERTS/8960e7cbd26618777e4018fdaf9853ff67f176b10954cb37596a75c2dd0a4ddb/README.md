# STIR/SHAKEN CA Ecosystem Compliance

## Certificate Definitely SHAKEN 709J but not SHAKEN 0007

Tested At: 02 Nov 22 15:09 UTC\
Initial Validity Period: 1 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=Definitely SHAKEN 709J but not SHAKEN 0007, OU=Trolling division, O=Low Latency Communications LLC, L=Birmingham, ST=Alabama, C=US\
Issuer: CN=Peeringhub Inc SHAKEN Intermediate CA 2, OU=Certification Authorities, O=Peeringhub Inc, C=US\
Link: https://sketchy.gay/shaken/llc-cert-1.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIDZzCCAw2gAwIBAgIQTLnZY0ljdw3oGpfAHrWfmjAKBggqhkjOPQQDAjB8MQswCQYDVQQGEwJVUzEXMBUGA1UECgwOUGVlcmluZ2h1YiBJbmMxIjAgBgNVBAsMGUNlcnRpZmljYXRpb24gQXV0aG9yaXRpZXMxMDAuBgNVBAMMJ1BlZXJpbmdodWIgSW5jIFNIQUtFTiBJbnRlcm1lZGlhdGUgQ0EgMjAeFw0yMjExMDExOTE2NTlaFw0yMjExMDIxODEzMzNaMIGuMQswCQYDVQQGEwJVUzEQMA4GA1UECAwHQWxhYmFtYTETMBEGA1UEBwwKQmlybWluZ2hhbTEnMCUGA1UECgweTG93IExhdGVuY3kgQ29tbXVuaWNhdGlvbnMgTExDMRowGAYDVQQLDBFUcm9sbGluZyBkaXZpc2lvbjEzMDEGA1UEAwwqRGVmaW5pdGVseSBTSEFLRU4gNzA5SiBidXQgbm90IFNIQUtFTiAwMDA3MFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEFBa93U6KMI99Xg3P0PASsrMgFT34leh7iBT%2Be2qH1H9tXSg1TFYy8FAIh56h3YmbCfF3gxZTNPEeJWJEjAsy%2FaOCATwwggE4MA4GA1UdDwEB%2FwQEAwIHgDAMBgNVHRMBAf8EAjAAMB0GA1UdDgQWBBSGr38CyrvkydrypjGzkldyhS0RnTAfBgNVHSMEGDAWgBSuoXNRiClXEcoMqfSxCm5OuEtNBzAXBgNVHSAEEDAOMAwGCmCGSAGG%2FwkBAQEwFgYIKwYBBQUHARoECjAIoAYWBDcwOUowgaYGA1UdHwSBnjCBmzCBmKA6oDiGNmh0dHBzOi8vYXV0aGVudGljYXRlLWFwaS5pY29uZWN0aXYuY29tL2Rvd25sb2FkL3YxL2NybKJapFgwVjEUMBIGA1UEBwwLQnJpZGdld2F0ZXIxCzAJBgNVBAgMAk5KMRMwEQYDVQQDDApTVEktUEEgQ1JMMQswCQYDVQQGEwJVUzEPMA0GA1UECgwGU1RJLVBBMAoGCCqGSM49BAMCA0gAMEUCIQCd%2FCn6Te9yWx9DHwe7qq7C4U1kTSQ0nzpAmSMdABWp8gIgGBCB0Rpow%2F6KYP%2FDJDXKhIV6ydrFF712k5W8GPG3NYE%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_ambiguous_identifier](../../ISSUES/e_us_cp_ambiguous_identifier/README.md) | error | US_SHAKEN_CP | Names used in the STI certificates shall represent an unambiguous identifier for the SP Subject |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_atis_certificate_policies](../../ISSUES/e_atis_certificate_policies/README.md) | error | ATIS1000080 | STI certificate shall contain '2.16.840.1.114569.1.1.3' policy |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |


Generated: 02 Nov 22 15:10 UTC