# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 0172

Tested At: 04 Oct 24 15:48 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -675 day(s)\
Subject: CN=SHAKEN 0172, O=CallTower, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/0172/cd86ac8e-a3d6-4ea8-b839-f341f33b6ccd.pem

[View certificate details](https://x509.io/?cert=MIICyTCCAm%2BgAwIBAgIQRXzwZRwYW%2BErnsXN4RarWzAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjIxMTIyMDgyNTA3WhcNMjIxMTI5MDgyNTA2WjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJQ2FsbFRvd2VyMRQwEgYDVQQDEwtTSEFLRU4gMDE3MjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABAoexVwjQ3TjPg1W%2BBU9z%2FZih6RzM1z0a4mpZ%2BAy%2FFDTjTqgQqguScZd24VCgjyqzvCOOJq3kyu%2BBDaVTkVZG2WjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUSsywY3oZkDZLpQ%2FfUX4ueSwkibQwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDE3MjAKBggqhkjOPQQDAgNIADBFAiEAtHhVolizegkRPId5DS%2Fl3Y4RphgV%2B918cBN0phKvwi4CIAz7grEHYGq%2BuGMc6ydmSD2nrks3d0ER1CwQyZRCUnDZ)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 04 Oct 24 16:29 UTC