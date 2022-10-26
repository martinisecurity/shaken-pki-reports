# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate b4d26c19a9c0aef72e9326e6bcf76f925cc71b01
Tested At: 2022-10-26 21:00:36 +0000 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: 1 day(s)\
Subject: CN=SHAKEN 366G, OU=SHAKEN, O=USA Digital, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.clearip.com/e9497931-96b3-4ff8-8306-1b1273847a4d/67ccee23ba400de44b132f6c0f5081b6.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC7jCCApOgAwIBAgIQRAdYXQjo5HirwYiihJRlMTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjAyMDIzMzVaFw0yMjEwMjcyMDIzMzRaMEoxCzAJBgNVBAYTAlVTMRQwEgYDVQQKEwtVU0EgRGlnaXRhbDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzY2RzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAAPHkeL12izRPlAMjTgt6nupTevtmcoR6nZzIqxAnFLOF8sXAGs%2BTnDk6MAGxXDTOruRiMIuXwrEWBuxg9rpsyjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUsTeIRIwGtetOJJQeubk6Q%2FfyGbIwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzY2RzAKBggqhkjOPQQDAgNJADBGAiEA4u8PRf3iC87serD3ZxwVv%2FrDhAIeAtFjO81jPjWi%2FuMCIQDTFRgIyi7kIhtXjRFpRvH9tKAESQ0VJj%2BFRaGKmb6Gug%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

Generated: 26/10/2022 at 21:01:13