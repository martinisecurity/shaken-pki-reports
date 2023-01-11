# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 1797 certificates were included in the corpus being tested
- 184 certificates in the corpus were skipped because they are duplicates
- 1250 certificates in the corpus were skipped because they are expired
- 5 certificates in the corpus were skipped because they are not currently trusted
- 358 certificates being tested against the remaining rules
- 4.09 issues on average found in unexpired, trusted, and non-compliant certificates
- 93.02% of certificates contain one or more Error level issue
- 50.56% of certificates contain one or more Warning level issue
- 1.12% of certificates contain one or more Notice level issue
- 17.32% of certificates are too old to be assessed against currently enforced expectations
- 361 days is the average remaining validity for the certificates in the corpus
- 362 days is the average initial validity for the certificates in the corpus
- 140 certificates expire in the next 30 days
- 4.98 average number of unexpired certificates per OCN observed
- 361 unique OCNs observed in unexpired and valid certificate corpus

### CA Certificates

- 38 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 1 certificates in the corpus were skipped because they are expired
- 4 certificates in the corpus were skipped because they are not currently trusted
- 33 certificates being tested against the remaining rules
- 2.05 issues on average found in unexpired, trusted, and non-compliant certificates
- 51.52% of certificates contain one or more Error level issue
- 51.52% of certificates contain one or more Warning level issue
- 3.03% of certificates contain one or more Notice level issue
- 69.70% of certificates are too old to be assessed against currently enforced expectations
- 5680 days is the average remaining validity for the certificates in the corpus
- 5612 days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository URL

- 57.82% of certificate repository URLs contain one or more Error level issue
- 85.47% of certificates repository URLs contain one or more Warning level issue
- 0.00% of certificates repository URLs contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#leaf-certificates) | 25 (6.98%) | 25 (100.00%) | 25 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#leaf-certificates) | 2 (0.56%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#leaf-certificates) | 24 (6.70%) | 0 (0.00%) | 0 (0.00%) | 1 (4.17%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#leaf-certificates) | 37 (10.34%) | 37 (100.00%) | 1 (2.70%) | 0 (0.00%) | 34 (91.89%) |
| [NetNumber](CERTS/NetNumber/README.md#leaf-certificates) | 3 (0.84%) | 3 (100.00%) | 0 (0.00%) | 3 (100.00%) | 0 (0.00%) |
| [Neustar](CERTS/Neustar/README.md#leaf-certificates) | 116 (32.40%) | 115 (99.14%) | 29 (25.00%) | 0 (0.00%) | 27 (23.28%) |
| [Peeringhub](CERTS/Peeringhub/README.md#leaf-certificates) | 9 (2.51%) | 9 (100.00%) | 1 (11.11%) | 0 (0.00%) | 0 (0.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#leaf-certificates) | 6 (1.68%) | 6 (100.00%) | 6 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [Sansay](CERTS/Sansay/README.md#leaf-certificates) | 105 (29.33%) | 105 (100.00%) | 105 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#leaf-certificates) | 2 (0.56%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) |
| [TransNexus](CERTS/TransNexus/README.md#leaf-certificates) | 29 (8.10%) | 29 (100.00%) | 12 (41.38%) | 0 (0.00%) | 1 (3.45%) |
| **Total** | 358 (100.00%) | 333 (93.02%) | 181 (50.56%) | 4 (1.12%) | 62 (17.32%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| [Comcast](CERTS/Comcast/README.md#ca-certificates) | 2 (6.06%) | 0 (0.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [GBSDTech](CERTS/GBSDTech/README.md#ca-certificates) | 2 (6.06%) | 1 (50.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [Martini Security](CERTS/Martini_Security/README.md#ca-certificates) | 3 (9.09%) | 0 (0.00%) | 0 (0.00%) | 1 (33.33%) | 0 (0.00%) |
| [Metaswitch](CERTS/Metaswitch/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 0 (0.00%) | 0 (0.00%) | 2 (100.00%) |
| [NetNumber](CERTS/NetNumber/README.md#ca-certificates) | 3 (9.09%) | 2 (66.67%) | 2 (66.67%) | 0 (0.00%) | 3 (100.00%) |
| [Neustar](CERTS/Neustar/README.md#ca-certificates) | 6 (18.18%) | 4 (66.67%) | 6 (100.00%) | 0 (0.00%) | 4 (66.67%) |
| [Peeringhub](CERTS/Peeringhub/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [Ribbon Communications](CERTS/Ribbon_Communications/README.md#ca-certificates) | 2 (6.06%) | 2 (100.00%) | 2 (100.00%) | 0 (0.00%) | 2 (100.00%) |
| [Sansay](CERTS/Sansay/README.md#ca-certificates) | 2 (6.06%) | 0 (0.00%) | 2 (100.00%) | 0 (0.00%) | 1 (50.00%) |
| [T-Mobile](CERTS/T-Mobile/README.md#ca-certificates) | 4 (12.12%) | 1 (25.00%) | 0 (0.00%) | 0 (0.00%) | 3 (75.00%) |
| [TransNexus](CERTS/TransNexus/README.md#ca-certificates) | 5 (15.15%) | 3 (60.00%) | 3 (60.00%) | 0 (0.00%) | 3 (60.00%) |
| **Total** | 33 (100.00%) | 17 (51.52%) | 17 (51.52%) | 1 (3.03%) | 23 (69.70%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 11 Jan 23 21:59 UTC