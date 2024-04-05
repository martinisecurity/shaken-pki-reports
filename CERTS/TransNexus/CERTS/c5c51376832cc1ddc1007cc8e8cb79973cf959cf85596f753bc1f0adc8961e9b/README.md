# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 148K

Tested At: 05 Apr 24 18:45 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -554 day(s)\
Subject: CN=SHAKEN 148K, OU=SHAKEN, O=Orange County Rural Electric Membership Coop, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/148K/340eb5b1-b461-4f5e-81f7-6da527199eb9.pem

[View certificate details](https://x509.io/?cert=MIIDDTCCArSgAwIBAgIQZf2fHIdoflcF%2BAljLpMrDzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA3MDExNTMwNDlaFw0yMjA5MjkxNTMwNDhaMGsxCzAJBgNVBAYTAlVTMTUwMwYDVQQKEyxPcmFuZ2UgQ291bnR5IFJ1cmFsIEVsZWN0cmljIE1lbWJlcnNoaXAgQ29vcDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMTQ4SzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABPF%2BR0eiP%2Bw4JVRRz%2BmZJG6WqzNsmsnY8U5LD9xJRDp88tQ8txLfG%2FUnVIuLCmPz6kk5PC7QjXguZvW0kpSPORSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUCab19HQtEStJySrqBcZMW74TGfwwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMTQ4SzAKBggqhkjOPQQDAgNHADBEAiBaEBLHM7JRTDcPV5yVFbpif0s9RtqAxVXtocvBmU%2ByvQIgPcPDZAGknZFnvwZdda2LX4KFkuyUIyHP7WCXPWtvE6c%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 05 Apr 24 19:04 UTC