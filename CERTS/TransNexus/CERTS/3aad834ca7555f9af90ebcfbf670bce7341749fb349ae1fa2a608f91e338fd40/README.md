# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 459J

Tested At: 04 Oct 24 15:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -722 day(s)\
Subject: CN=SHAKEN 459J, OU=SHAKEN, O=Altaworx, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/d3d6dbfc-2914-49c7-8f47-d0aa5bd5d764/666215ba7818d320f99063ea96558fd4.pem

[View certificate details](https://x509.io/?cert=MIIC6jCCApCgAwIBAgIQQgymwtBPjIYOyw1iWaIDXTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDUyMDIxMTdaFw0yMjEwMTIyMDIxMTZaMEcxCzAJBgNVBAYTAlVTMREwDwYDVQQKEwhBbHRhd29yeDEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gNDU5SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAEinjolbLh1%2FKMXQiGEWMg7x7HxbRcJrwljw8Mgfj9vDss%2BnUTJBwF%2BAsjOIZAHK30NOR4yF%2BZA5%2BGt0O9tp%2FijggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUzw2m%2FfbeUdbUKX%2BxnY1uy61ZisMwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENDU5SjAKBggqhkjOPQQDAgNIADBFAiBpX6HsySM7ArMwLtlmRcHqoaBIWB27fZdrMvMv0hTbsgIhAPM43qxEt7mljCrOnGKMsmel3jHUS0mWRLYYTnqPL8Ab)

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