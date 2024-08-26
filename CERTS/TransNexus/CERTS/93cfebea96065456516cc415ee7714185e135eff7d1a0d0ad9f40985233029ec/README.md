# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 157C

Tested At: 26 Aug 24 18:06 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -679 day(s)\
Subject: CN=SHAKEN 157C, OU=SHAKEN, O=DigitalSpeed Communications, C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/82d97ecb-9b38-4926-9062-8e21b8986930/28ec7e97780e7ae26bbe623cc083537c.pem

[View certificate details](https://x509.io/?cert=MIIC%2FDCCAqOgAwIBAgIQdaI53SlqzpVwRdlOOTskcDAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMDkyMDUwNDdaFw0yMjEwMTYyMDUwNDZaMFoxCzAJBgNVBAYTAlVTMSQwIgYDVQQKExtEaWdpdGFsU3BlZWQgQ29tbXVuaWNhdGlvbnMxDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE1N0MwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQiIM7JktGOqw%2B9g0V23rXN%2BMYme7q8FV%2BjUeWU8Sti9tbSJB9d82Vgs7FyLA0PzWhTvPl%2BnyTnOQDhcK9b6KbVo4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFNA6BTX0lNpg34GFY4ZgDCNEynKBMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE1N0MwCgYIKoZIzj0EAwIDRwAwRAIgOo%2Fwl6QBJYuPa%2F0GTzkD81hMMsH8jAURvHPqF3TdQGwCIAVAvQiJNJpRYcPPVyqNQtsufJ3Tfw92Wysyz1VwYs46)

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


Generated: 26 Aug 24 18:49 UTC