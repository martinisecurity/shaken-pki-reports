# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 551G

Tested At: 05 Apr 24 18:45 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -500 day(s)\
Subject: CN=SHAKEN 551G, O=Brightlink Communications LLC, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/551G/80afad74-7d8f-4c50-a4cf-903042b30608.pem

[View certificate details](https://x509.io/?cert=MIIC7zCCApSgAwIBAgIQS8z1qJz1dJcU%2FjUSjNaL%2FTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjExMTUxMjQyMzVaFw0yMjExMjIxMjQyMzRaMEsxCzAJBgNVBAYTAlVTMSYwJAYDVQQKEx1CcmlnaHRsaW5rIENvbW11bmljYXRpb25zIExMQzEUMBIGA1UEAxMLU0hBS0VOIDU1MUcwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAARv4hIxEolc7X0j4B7qlxK098DbWDko0jv2jOj%2BWzcRXU3ar5kXoa0h76SVCo7rcXeo6stfyCHIjmEtUjF%2BJHDPo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCB4AwHQYDVR0OBBYEFNrQd8lwubixV8Fiyz97zK4Xv7j1MB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDU1MUcwCgYIKoZIzj0EAwIDSQAwRgIhAIpHeNkZ1mwTaYg2WUTLNRnZtiB6GZAEK1NgDxkklzmDAiEA2wbLREDbbkVnGyUznIf2zf1mlgFe3vc0P%2FMLihRXZzo%3D)

| Code | Type | Source | Details |
|------|------|--------|---------|
| [e_us_cp_subject_sn_shall](../../ISSUES/e_us_cp_subject_sn_shall/README.md) | error | US_SHAKEN_CP | The DN does not contain a serialNumber attribute. |

### Not Effective

- e_atis_ext_not_specified
- e_atis_serial_number_size
- e_atis_subject_c_iso
- e_atis_tn_auth_list_spc_format

\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.


Generated: 05 Apr 24 19:04 UTC