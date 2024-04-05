# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 505J

Tested At: 05 Apr 24 18:45 UTC\
Initial Validity Period: 90 day(s)\
Remaining Validity Period: -467 day(s)\
Subject: CN=SHAKEN 505J, OU=SHAKEN, O=HFA Services LLC dba Call48, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.transnexus.com/505J/e7c8b693-b40a-4cfd-bd50-260db797f09b.pem

[View certificate details](https://x509.io/?cert=MIIC%2FDCCAqOgAwIBAgIQQ%2B4HO%2BQ0PMQ%2BTQqxY5JcsTAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjA5MjYxODQwMjBaFw0yMjEyMjUxODQwMTlaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtIRkEgU2VydmljZXMgTExDIGRiYSBDYWxsNDgxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDUwNUowWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAASlXtexQZmHx3Fp%2FizBV%2F8MA1CW6S%2Bfkq1iQMJqxW8ATiBUP8qIOcOwG5xLRFFooMJmhgWL1lR%2FQ8fkVLQUjWLqo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFGN%2BeoA7Pzi5oTtCzKDgKpjX%2F00nMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDUwNUowCgYIKoZIzj0EAwIDRwAwRAIgMMTUp0wkYfBuL70U6aocEI1DeAkknvlGk3%2B0N6%2B262kCID4dFx4T9ZfvRODsx5gK4615PnDjM7fpx%2FR0P3gthcDb)

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


Generated: 05 Apr 24 19:04 UTC