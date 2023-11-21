# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 130B

Tested At: 21 Nov 23 17:38 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -398 day(s)\
Subject: CN=SHAKEN 130B, OU=SHAKEN, O=Global Telecom, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4efae05c-c8cd-405b-ae97-c62a52afd074/2f9fbd9cb49279c3b553c6bd9ee383ee.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC7zCCApagAwIBAgIQZSMF5au479ZUnQ2ZM3iS2TAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTIxMjEwNTJaFw0yMjEwMTkxMjEwNTFaME0xCzAJBgNVBAYTAlVTMRcwFQYDVQQKEw5HbG9iYWwgVGVsZWNvbTEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMTMwQjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABJqhX%2FfSJY0JUuRuj8tvGu2apM%2BLyXbOaC8wLcv9FzQjBwzFnH44WRkhL7evrWgqMV5XwL7zbptNrOY5WGgi%2FC6jggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUjR0Zxagta0mpWbgQoNLKUj9lK4QwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMTMwQjAKBggqhkjOPQQDAgNHADBEAiANL1RmkjhNXfoIeKEGWpTwI%2F8fOL7ZziPV95uK6X2i0AIgIgQVf8Cx%2B0rLueLMyFaBLGTdxvUl09sr%2BGYu04PV%2F9E%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 17:53 UTC