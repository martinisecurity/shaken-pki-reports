# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 769J

Tested At: 21 Nov 23 16:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -398 day(s)\
Subject: CN=SHAKEN 769J, OU=SHAKEN, O=Affiliated Technology Solutions LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/471a46b3-3de4-4aa2-964b-2ff03ddbc8fe/5399aa88e45dc2273bf83cc11b18b3a7.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIDBTCCAqugAwIBAgIQax9PzB77yOJVE2PlTOrTqzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMDIwMTNaFw0yMjEwMTgyMDIwMTJaMGIxCzAJBgNVBAYTAlVTMSwwKgYDVQQKEyNBZmZpbGlhdGVkIFRlY2hub2xvZ3kgU29sdXRpb25zIExMQzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNzY5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDv1EVQNI7dSoMWy%2BdxU97yY2DekRJ24Q%2BoF0bJsSHuyJMbK9LnlYRzYcarvHQM33U39vobiUaWAnnSYRvcYD9ijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQU1L1idwQohjzwYqSgV1zD7lR1vaAwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzY5SjAKBggqhkjOPQQDAgNIADBFAiEAl16Vh31BbxQA6gSyC9%2BILHlGhthzqvhV3QF47yk0nHoCIAxF5NeKG39GLQAXCmXlYWbb%2FNrRZ2opnCWJOPmg56pR)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:16 UTC