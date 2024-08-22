# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 918J

Tested At: 22 Aug 24 15:20 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -699 day(s)\
Subject: CN=SHAKEN 918J, OU=SHAKEN, O=VoIPX International Inc, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/3bbe1f4c-3184-44ee-84f4-2f63891aa57b/d9b86d264d8bfdcea8367f36c90b1ee6.pem

[View certificate details](https://x509.io/?cert=MIIC%2BTCCAp%2BgAwIBAgIQXWUz7ha2JdOPcOU9IqtIsTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MTUxOTQxMjhaFw0yMjA5MjIxOTQxMjdaMFYxCzAJBgNVBAYTAlVTMSAwHgYDVQQKExdWb0lQWCBJbnRlcm5hdGlvbmFsIEluYzEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gOTE4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABGKx%2B8FLgWyoVcVXSHpCaq%2FN6Tn3KmAFp4cziuo9AQN5A8WR%2BigRk37bxwfaBqnEb1%2BVmCqf%2B%2F74G8U15LQz6mKjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUaKT4l67xfF3j83%2FGZRYlz3qX2jUwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEOTE4SjAKBggqhkjOPQQDAgNIADBFAiB3CG2xtoF61duhLwO9p46XejMbES84QYe7cMrh9UclUwIhAKKfySQ2Rpgw3%2FZmNmP%2B3ddgXapagFxmPCjaercsqj9y)

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


Generated: 22 Aug 24 16:06 UTC