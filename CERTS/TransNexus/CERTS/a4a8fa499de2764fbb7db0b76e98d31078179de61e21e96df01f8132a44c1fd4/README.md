# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 012K

Tested At: 22 Aug 24 15:19 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -659 day(s)\
Subject: CN=SHAKEN 012K, OU=SHAKEN, O=CallCurrent\\, Inc., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/06ebd24a-1f0a-46d5-8a2f-a7ae49be8ed6/2d9c4131825627aaf47bd2c9471d43ae.pem

[View certificate details](https://x509.io/?cert=MIIC8jCCApmgAwIBAgIQfhaVwnud0cgniW3bhwRLtDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMjYwODM5MTZaFw0yMjExMDIwODM5MTVaMFAxCzAJBgNVBAYTAlVTMRowGAYDVQQKExFDYWxsQ3VycmVudCwgSW5jLjEPMA0GA1UECxMGU0hBS0VOMRQwEgYDVQQDEwtTSEFLRU4gMDEySzBZMBMGByqGSM49AgEGCCqGSM49AwEHA0IABO6nm%2FoHAJKtfAf2GenWojeiWVrzisFDY9Wmt%2FedB9rGIC9BcUxTmn%2BGp5npDuUWC%2Fg0Pq8bzBOAHgFXm8GO1ayjggE8MIIBODAMBgNVHRMBAf8EAjAAMA4GA1UdDwEB%2FwQEAwIAgDAdBgNVHQ4EFgQUuokK9EtcLDfavkpp10kP7etClJEwHwYDVR0jBBgwFoAUu5beMRLN05aZhKQ2MGA811KBfScwFwYDVR0gBBAwDjAMBgpghkgBhv8JAQEDMIGmBgNVHR8EgZ4wgZswgZigOqA4hjZodHRwczovL2F1dGhlbnRpY2F0ZS1hcGkuaWNvbmVjdGl2LmNvbS9kb3dubG9hZC92MS9jcmyiWqRYMFYxFDASBgNVBAcMC0JyaWRnZXdhdGVyMQswCQYDVQQIDAJOSjETMBEGA1UEAwwKU1RJLVBBIENSTDELMAkGA1UEBhMCVVMxDzANBgNVBAoMBlNUSS1QQTAWBggrBgEFBQcBGgQKMAigBhYEMDEySzAKBggqhkjOPQQDAgNHADBEAiBRb697byLXj888AHfwG8QqIcUkzPA%2FZHBHdcC%2BdWV8AQIgNwPyq8%2BWqqaJjnOZAYZcqTtffJRgkGEgj5MJ2r5d%2BeQ%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_incorrect_ku_encoding](../../ISSUES/e_incorrect_ku_encoding/README.md) | error | RFC5280 | KeyUsage contains an inefficient encoding wherein the number of 'unused bits' is declared to be 0, but it should be 7. Raw Bytes: [3 2 0 128], Raw Binary: [00000011 00000010 00000000 10000000] |
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 22 Aug 24 16:06 UTC