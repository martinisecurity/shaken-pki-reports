# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 8877 certificates were included in the corpus being tested
- 724 certificates in the corpus were skipped because they are duplicates
- 6836 certificates in the corpus were skipped because they are expired
- 474 certificates in the corpus were skipped because they are not currently trusted
- 843 certificates being tested against the remaining rules
- 1.82 issues on average found in unexpired, trusted, and non-compliant certificates
- 55.40% of certificates contain one or more Error level issue
- 1.42% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 13.76% of certificates are too old to be assessed against currently enforced expectations
- 279 days is the average remaining validity for the certificates in the corpus
- 279 days is the average initial validity for the certificates in the corpus
- 255 certificates expire in the next 30 days
- 13.51 average number of unexpired certificates per OCN observed
- 657 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 44 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 8 certificates in the corpus were skipped because they are not currently trusted
- 36 certificates being tested against the remaining rules
- 2.60 issues on average found in unexpired, trusted, and non-compliant certificates
- 41.67% of certificates contain one or more Error level issue
- 0.00% of certificates contain one or more Warning level issue
- 0.00% of certificates contain one or more Notice level issue
- 80.56% of certificates are too old to be assessed against currently enforced expectations
- 5751 days is the average remaining validity for the certificates in the corpus
- 5661 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 36.54% of certificate repository URLs contain one or more Error level issue
- 53.50% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 12 (1.42%) | 12 (100.00%) | 12 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 3 (0.36%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 371 (44.01%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 1 (0.27%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 71 (8.42%) | 71 (100.00%) | 0 (0.00%) | 0 (0.00%) | 50 (70.42%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 5 (0.59%) | 5 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 150 (17.79%) | 150 (100.00%) | 0 (0.00%) | 0 (0.00%) | 47 (31.33%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 17 (2.02%) | 17 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 25 (2.97%) | 25 (100.00%) | 0 (0.00%) | 0 (0.00%) | 4 (16.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 145 (17.20%) | 145 (100.00%) | 0 (0.00%) | 0 (0.00%) | 11 (7.59%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 1 (0.12%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Telonium](CERTS/Telonium/README.md#leaf-certificates) | 20 (2.37%) | 15 (75.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 23 (2.73%) | 23 (100.00%) | 0 (0.00%) | 0 (0.00%) | 3 (13.04%) |
| **Total** | 843 (100.00%) | 467 (55.40%) | 12 (1.42%) | 0 (0.00%) | 116 (13.76%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [CTIA](CERTS/CTIA/README.md#ca-certificates) | 1 (2.78%) | 1 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (5.56%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (5.56%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 2 (5.56%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 3 (8.33%) | 3 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (66.67%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (8.33%) | 2 (66.67%) | 0 (0.00%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (16.67%) | 1 (16.67%) | 0 (0.00%) | 0 (0.00%) | 6 (100.00%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (5.56%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (5.56%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (5.56%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (11.11%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 4 (100.00%) |
| [Telonium](CERTS/Telonium/README.md#ca-certificates) | 5 (13.89%) | 3 (60.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 2 (5.56%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| **Total** | 36 (100.00%) | 15 (41.67%) | 0 (0.00%) | 0 (0.00%) | 29 (80.56%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 21 Nov 23 17:16 UTC