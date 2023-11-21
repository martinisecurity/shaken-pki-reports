# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 738J

Tested At: 21 Nov 23 01:30 UTC\
Initial Validity Period: 9 day(s)\
Remaining Validity Period: -285 day(s)\
Subject: CN=SHAKEN 738J, O=SkySwitch, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA4, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/738J/a91a00f6-3202-4e8e-ba72-2f69e2436525.pem

[View certificate details](https://understandingwebpki.com/?cert=MIICyTCCAm%2BgAwIBAgIQbT6VWezqqpTFFR%2BvyOnnXTAKBggqhkjOPQQDAjBWMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEsMCoGA1UEAxMjVHJhbnNOZXh1cywgSW5jLiBTSEFLRU4gSXNzdWluZyBDQTQwHhcNMjMwMTMwMjAzNTA1WhcNMjMwMjA4MjAzNTA0WjA3MQswCQYDVQQGEwJVUzESMBAGA1UEChMJU2t5U3dpdGNoMRQwEgYDVQQDEwtTSEFLRU4gNzM4SjBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABIur40jaQMTWIgA2K50eaMSNFceeojJb5eeuQfhyAovIsStyT%2BD8jxpjcGQtBJWQo1PtoNuBZ7apkw3XhVnPAhGjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIHgDAdBgNVHQ4EFgQUFsVdySLMf8NzxDQTGBBZNUW9qVQwHwYDVR0jBBgwFoAUMPX18rfhSwsPEHV9SKSQ8T2LKCowFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYENzM4SjAKBggqhkjOPQQDAgNIADBFAiBnuJwOnAle512UBY2vJTKr%2B%2BtDrUyRbgFauILIk7z8wwIhAKjXtNUkXPVCeRv43uol3GgwgPJ63QvImyhmlmoOl2sw)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 21 Nov 23 01:55 UTC