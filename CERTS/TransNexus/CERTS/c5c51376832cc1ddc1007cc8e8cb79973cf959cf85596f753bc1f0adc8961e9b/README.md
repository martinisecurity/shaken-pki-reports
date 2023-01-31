# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 148K

Tested At: 31 Jan 23 21:39 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -124 day(s)\
Subject: CN=SHAKEN 148K, OU=SHAKEN, O=Orange County Rural Electric Membership Coop, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/148K/340eb5b1-b461-4f5e-81f7-6da527199eb9.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDDTCCArSgAwIBAgIQZf2fHIdoflcF%2BAljLpMrDzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDExNTMwNDlaFw0yMjA5MjkxNTMwNDhaMGsxCzAJBgNVBAYTAlVTMTUwMwYDVQQKEyxPcmFuZ2UgQ291bnR5IFJ1cmFsIEVsZWN0cmljIE1lbWJlcnNoaXAgQ29vcDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMTQ4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPF%2BR0eiP%2Bw4JVRRz%2BmZJG6WqzNsmsnY8U5LD9xJRDp88tQ8txLfG%2FUnVIuLCmPz6kk5PC7QjXguZvW0kpSPORSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUCab19HQtEStJySrqBcZMW74TGfwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMTQ4SzAKBggqhkjOPQQDAgNHADBEAiBaEBLHM7JRTDcPV5yVFbpif0s9RtqAxVXtocvBmU%2ByvQIgPcPDZAGknZFnvwZdda2LX4KFkuyUIyHP7WCXPWtvE6c%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |


Generated: 31 Jan 23 21:50 UTC