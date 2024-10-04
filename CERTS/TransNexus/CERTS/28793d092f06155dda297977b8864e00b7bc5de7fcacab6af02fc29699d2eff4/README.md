# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 625J

Tested At: 04 Oct 24 15:43 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -713 day(s)\
Subject: CN=SHAKEN 625J, OU=SHAKEN, O=Victory Telecom Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/b7343686-5ed8-402c-89a3-8bf1a3d48975/61649b0719ddf13f46df03bd3ca7845a.pem

[View certificate details](https://x509.io/?cert=MIIC9jCCApygAwIBAgIQdS8CcPOQiNhsr849%2BKQQjzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTQyMDIxNDhaFw0yMjEwMjEyMDIxNDdaMFMxCzAJBgNVBAYTAlVTMR0wGwYDVQQKExRWaWN0b3J5IFRlbGVjb20gSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNjI1SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAH7j2riWgHAwgoYc9OUn4HsYWm42ewlOcEZNT%2B25nwf%2F5EOda4oGLyZFG1ROHQEL8z7nPWNWfNkTUuNRHs0Un%2BjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQURMQdJjlwcssLL9kQyOtqYaVg7BcwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENjI1SjAKBggqhkjOPQQDAgNIADBFAiA48Qe9w0gdKbF7XOFpqIFeU9POAOY2cKUJKOYq9g7VKgIhAK0mUXlTuKSaa%2FAcTTtS0F4HIyODgdhTXr81qZn1GgjQ)

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


Generated: 04 Oct 24 16:29 UTC