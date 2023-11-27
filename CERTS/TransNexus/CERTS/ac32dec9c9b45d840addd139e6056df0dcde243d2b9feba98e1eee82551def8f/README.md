# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0226

Tested At: 27 Nov 23 22:25 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -429 day(s)\
Subject: CN=SHAKEN 0226, OU=SHAKEN, O=Lumos Networks, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/f8a68e44-8fb0-4f28-b533-c4df27ed8e1b/0b76cb34c9170502d9a00a59220461a4.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8TCCApagAwIBAgIQdD5zKcgG%2BZlG3bHKJvGvSDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTcyMDIwMDlaFw0yMjA5MjQyMDIwMDhaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5MdW1vcyBOZXR3b3JrczEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDIyNjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABDXnhGC9EknS3WEyrjmhYdFf2ocmZVCqzPi3dBm%2BVXL9UKITL8gd1P2Uws5g%2FcsEVjJl29q3qZbyoHIYIFWrymmjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUpNZdXL3fgRyc6UE9FNpUe2UhN%2BgwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDIyNjAKBggqhkjOPQQDAgNJADBGAiEAt4tef67Yer9hwTSBDEUFio%2BmxP1LV31IEtwVRzucpIUCIQC7XmJYzXgHJaYJItnhSjsDUlN5x3EbPbBuTSUDvjLtIA%3D%3D)

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


Generated: 27 Nov 23 22:56 UTC