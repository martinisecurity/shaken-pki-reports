# STIR/SHAKEN CA Ecosystem Compliance

[Approved Certificate Authorities](https://ecosystemcompliance.martinisecurity.com/#:~:text=Approved%20Certificate%20Authorities) in the STIR/SHAKEN ecosystem are required to meet technical requirements from [ATIS-1000080](https://access.atis.org/apps/group_public/document.php?document_id=62163) and policy requirements from the supporting CA ecosystem’s [Certificate Policy](https://authenticate.iconectiv.com/documents-authenticate).

This report is broken int two parts:
1. One generated using [Zlint](https://github.com/zmap/zlint) a tool commonly used to asses CA ecosystem compliance with such requirements. The tests used to generate this report are currently not part of the main [Zlint](https://github.com/martinisecurity/zlint) distribution but can be found here.
2. One generated with a custom script that eumerates the known STIR/SHAKEN certificates and asses each repository against the current rule set . The source for this test can be found here while the report itself can be found [here](REPOS/README.md).

## Summary

### Leaf Certificates

- 0 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 0 certificates being tested against the remaining rules
- NaN% of certificates contain one or more Error level issue
- NaN% of certificates contain one or more Warning level issue
- NaN% of certificates contain one or more Notice level issue
- NaN% of certificates are too old to be assessed against currently enforced expectations
- NaN days is the average remaining validity for the certificates in the corpus
- NaN days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

### CA Certificates

- 0 certificates were included in the corpus being tested
- 0 certificates in the corpus were skipped because they are duplicates
- 0 certificates in the corpus were skipped because they are expired
- 0 certificates in the corpus were skipped because they are not currently trusted
- 0 certificates being tested against the remaining rules
- NaN% of certificates contain one or more Error level issue
- NaN% of certificates contain one or more Warning level issue
- NaN% of certificates contain one or more Notice level issue
- NaN% of certificates are too old to be assessed against currently enforced expectations
- NaN days is the average remaining validity for the certificates in the corpus
- NaN days is the average initial validity for the certificates in the corpus
- 0 certificates expire in the next 30 days

## Certificate Repository

- NaN% of certificate repositories contain one or more Error level issue
- NaN% of certificates repositories contain one or more Warning level issue
- NaN% of certificates repositories contain one or more Notice level issue

## Details

\* The percent of certificates per issuer is calculated against total certificates from all issuers.\
\*\* The percent of errors, warnings and notices is calculated against total observed certificates from the specified issuer.\
\*\*\* Tests use the ATIS-1000080 and Certificate Policy versions release dates to determine if tests are ran. Certificates issued before these dates are not executed as the rules may not have been enforce at the time.

### Leaf Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| **Total** | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) |

### CA Certificates

| Issuers | Certificates | Errors | Warnings | Notices | Not Effective |
|---------|--------------|--------|----------|---------|---------------|
| **Total** | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) | 0 (NaN%) |

### Key

| Type | Description |
|------|-------------|
| Errors | Tests in which the specifications are unambiguous on what the expected behavior must be. |
| Warnings | Tests in which the specifications are ambiguous or are provide only a recommendation. |
| Notices | Tests in which industry best practices are not followed. |
| Not Effective | Tests that exist in the current specifications but were not in effect at the time of issuance. |


Generated: 01/11/2022 at 16:11:23