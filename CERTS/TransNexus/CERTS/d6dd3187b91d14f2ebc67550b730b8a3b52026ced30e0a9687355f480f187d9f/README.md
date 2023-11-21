# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 459J

Tested At: 21 Nov 23 17:39 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -398 day(s)\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/3caad53466f33f917b5c468042ec6609.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC6TCCApCgAwIBAgIQQ7CVyGPZMQ9Yok%2FW040h3jAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIxMzhaFw0yMjEwMTgyMDIxMzdaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABKtVGQqsidprYL2WWjmVLa7WFn%2F0gkX1x9mpJXVj%2Bi4G%2BfIBdN8kev2MFOHZK4F9Ke7OKxpOkSWLJnd9cOoHVvujggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUlp9ieVZQifozhaTUAfeiGVxnOBgwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNHADBEAiBBHyGRlAFrQagyBNlkyh6BhttwyZooMCeNPv3MD1HNQAIgB%2FI59aWLLiwgAsL17DGp55LmjAG3XveSrLaqZQiyZEc%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC