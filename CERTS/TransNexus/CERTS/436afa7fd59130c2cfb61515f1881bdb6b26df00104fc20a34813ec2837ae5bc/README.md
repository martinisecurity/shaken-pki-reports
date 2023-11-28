# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 625J

Tested At: 28 Nov 23 10:24 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -390 day(s)\
Subject: CN=SHAKEN 625J, OU=SHAKEN, O=Victory Telecom Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b7343686-5ed8-402c-89a3-8bf1a3d48975/1b6c455dcc3cd7586f99609b4a7291af.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC9jCCApygAwIBAgIQamaEhhL49kHAMJ7kudCLCDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYyMDIzNTJaFw0yMjExMDIyMDIzNTFaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRWaWN0b3J5IFRlbGVjb20gSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNjI1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABMnDPvAXs57ZPLTqSu2fp7yO95QjrIitRO4WFWR7g3r1CxDUrjqrh%2FtIjogdtT4u4lvcDnp%2FZqT7J6fqaY3vFnijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUSpTShpUvU4p0uSEJbRjFqfo%2BrSkwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENjI1SjAKBggqhkjOPQQDAgNIADBFAiB2o%2F5lNDsmK4xGni5sMogw5sQUhpFpjIg19Af%2F5i6IyAIhALDMC3YMNY8AtGrXQK50zLjk4I8pbvaVdNP%2Bhu9YUIUh)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 28 Nov 23 10:53 UTC