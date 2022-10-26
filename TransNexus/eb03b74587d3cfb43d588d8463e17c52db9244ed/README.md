# STIR/SHAKEN CA Ecosystem Compliance
## TransNexus

### Certificate eb03b74587d3cfb43d588d8463e17c52db9244ed
Tested At: 2022-10-26 21:01:03 +0000 UTC\
Initial Validity Period: 30 day(s)\
Remaining Validity Period: 5 day(s)\
Subject: CN=SHAKEN 345J, OU=SHAKEN, O=Ooma Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US

Link: https://certificates.transnexus.com/345J/31d44cfd-02e8-4a99-830b-c56a58de1835.pem

View: [Click to view](https://understandingwebpki.com/?cert=MIIC6zCCApCgAwIBAgIQVAZVDoxUtd2Z8ozIQnEHbDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDEwNDA0MTZaFw0yMjEwMzEwNDA0MTVaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhPb21hIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMzQ1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABFiy9OeR8egLCKfIVyPRs1UftGsoCE6RyNxRwdQ9PHKJx9vKHpsNoNVvKOUaiZh1tKGrHd4YO46TlkXrCHN62kejggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUIJY67HjeQJEZlhimvicpeVqGbicwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMzQ1SjAKBggqhkjOPQQDAgNJADBGAiEAnk%2F1XxBSYAOt4HIbL3C6E0O0BhwfdsfzPQa6iWPbDuACIQDqXEd6VJp%2F4Es8R4gCIZ9SzzJZhxh8Y0Ge%2FbLECZAclA%3D%3D)


| Code | Type | Source | Details |
|------|------|--------|---------|
| w_cp1_3_subject_rdn_unknown | warn | CPv1.3 | Only CN, C, O, and SERIALNUMBER can be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| e_incorrect_ku_encoding | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| e_sti_crl_distribution | error | ATIS-1000080v4 | CRL Distribution Point shall be reachable if the requesting IP address within the program ACLs |
| e_cp1_3_subject_sn | error | CPv1.3 | STI certificate shall include a ‘serialNumber’ attribute along with the CN |

Generated: 26/10/2022 at 21:01:13