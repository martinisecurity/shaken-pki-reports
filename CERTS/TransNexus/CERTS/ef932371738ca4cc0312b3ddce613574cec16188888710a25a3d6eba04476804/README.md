# STIR/SHAKEN CA Ecosystem Compliance

## Certificate SHAKEN 1821

Tested At: 27 Nov 23 22:23 UTC\
Initial Validity Period: 7 day(s)\
Remaining Validity Period: -405 day(s)\
Subject: CN=SHAKEN 1821, OU=SHAKEN, O=InPhonex.com\\, LLC., C=US\
Issuer: CN=TransNexus\\, Inc. SHAKEN Issuing CA3, OU=SHAKEN, O=TransNexus\\, Inc., C=US\
Link: https://certificates.clearip.com/4eed3a39-8f41-4acb-9b86-42a558b3455a/a800a09771486ab5f1f6b2de2cffc00c.pem

[View certificate details](https://understandingwebpki.com/?cert=MIIC8zCCApqgAwIBAgIQfj4QO8zeuveVjabC1EZyRzAKBggqhkjOPQQDAjBnMQswCQYDVQQGEwJVUzEZMBcGA1UEChMQVHJhbnNOZXh1cywgSW5jLjEPMA0GA1UECxMGU0hBS0VOMSwwKgYDVQQDEyNUcmFuc05leHVzLCBJbmMuIFNIQUtFTiBJc3N1aW5nIENBMzAeFw0yMjEwMTEyMTA2MTVaFw0yMjEwMTgyMTA2MTRaMFExCzAJBgNVBAYTAlVTMRswGQYDVQQKExJJblBob25leC5jb20sIExMQy4xDzANBgNVBAsTBlNIQUtFTjEUMBIGA1UEAxMLU0hBS0VOIDE4MjEwWTATBgcqhkjOPQIBBggqhkjOPQMBBwNCAAQHUfti9Vkr6Uvl%2F0RXiXOTYkSLzYiVcE4gMMj%2FeDDJudH%2BiigOknG3PneyegQBOk0SSP%2FyLSU9FlSsIe5hdxJ3o4IBPDCCATgwDAYDVR0TAQH%2FBAIwADAOBgNVHQ8BAf8EBAMCAIAwHQYDVR0OBBYEFJoYvpeoAvYqv1vuv2EwIRrxfd1qMB8GA1UdIwQYMBaAFLuW3jESzdOWmYSkNjBgPNdSgX0nMBcGA1UdIAQQMA4wDAYKYIZIAYb%2FCQEBAzCBpgYDVR0fBIGeMIGbMIGYoDqgOIY2aHR0cHM6Ly9hdXRoZW50aWNhdGUtYXBpLmljb25lY3Rpdi5jb20vZG93bmxvYWQvdjEvY3JsolqkWDBWMRQwEgYDVQQHDAtCcmlkZ2V3YXRlcjELMAkGA1UECAwCTkoxEzARBgNVBAMMClNUSS1QQSBDUkwxCzAJBgNVBAYTAlVTMQ8wDQYDVQQKDAZTVEktUEEwFgYIKwYBBQUHARoECjAIoAYWBDE4MjEwCgYIKoZIzj0EAwIDRwAwRAIgfkGfhyAop9ex0uwLUFvW5aIyT2KZYviCLjJ%2BHavfwBkCIBwSZ8jXI96RuAlBxp%2BeVkek9q9HYahZcN651osLvUyT)

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


Generated: 27 Nov 23 22:56 UTC