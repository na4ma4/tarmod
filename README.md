# tarmod

Tool for helping make identical docker builds in CI by modifying the layers to reduce date/time differences.

Concept in progress.

## Concept

When building with docker on separate machines, the layers can contain identical data,
but the date/times on created directories will be based on build time, which makes the
layers different even though the data content is identical.

This tool should be given the tar file from `docker image save` and be able to recreate the layers with directory times based on the files contained and then recalculate the hashes for the layers so when `docker image import` is used it will result in identical images even on different machines so long as the content (and file dates) are identical.

A separate tool should be made to touch files with dates based on `git log` output.
