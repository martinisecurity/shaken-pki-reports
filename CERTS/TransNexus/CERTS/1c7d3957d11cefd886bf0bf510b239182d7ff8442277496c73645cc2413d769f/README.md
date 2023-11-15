# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 918J

Tested At: 15 Nov 23 16:12 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -388 day(s)\
Subject: CN=SHAKEN 918J, OU=SHAKEN, O=VoIPX International Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3bbe1f4c-3184-44ee-84f4-2f63891aa57b/a99ebf977abf4bb8bfe108293a6160c2.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC%2BDCCAp%2BgAwIBAgIQUNLkkESXOxBw53yycnAGMzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTUxOTQ0MjJaFw0yMjEwMjIxOTQ0MjFaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdWb0lQWCBJbnRlcm5hdGlvbmFsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABHIL85QDWytT2FxMmeEyq1sXPr79v%2FDe5Cgc05BT%2FVH711Dtd9muOdyL4s%2FMmjtiqmoXuhnHiHg0UAZVMD1HhCSjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUIOF4k2iEB1xq6GN4Omp0JfGFuZcwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTE4SjAKBggqhkjOPQQDAgNHADBEAiBlcLRgrxEQZ7%2BjJctj22mSPBjX9GJdHQIc3S9%2Fq7B9UQIgcwMUD6Q7%2FNQYiZRIdvW%2BVaZD02cKLl%2ByY7GjGwOwDOw%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [w_shaken_subject_rdn_unknown](../../ISSUES/w_shaken_subject_rdn_unknown/README.md) | warn | SHAKEN_PKI_BEST_PRACTICES | Only CN, C, O, L, and SERIALNUMBER should be included. Additional RNDs may introduce ambiguity and may not be verifiable |
| [e_us_cp_subject_sn](../../ISSUES/e_us_cp_subject_sn/README.md) | error | US_SHAKEN_CP | STI certificate shall include a ‘serialNumber’ attribute along with the CN |


Generated: 15 Nov 23 17:17 UTC