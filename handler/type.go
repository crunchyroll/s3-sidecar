package handler

import "time"

// CustomInput - will allow callers to pass custom input for setting up the
// header or the query-string for the underlying HTTP call
type CustomInput struct {
	// ** location: header

	// Return the object only if its entity tag (ETag) is the same as the one specified,
	// otherwise return a 412
	IfMatch *string

	// Return the object only if it has been modified since the specified time,
	// otherwise return a 304
	IfModifiedSince *time.Time

	// Return the object only if its entity tag (ETag) is different from the one specified
	// otherwise return a 304 (not modified).
	IfNoneMatch *string

	// Return the object only if it has not been modified since the specified time
	// otherwise return a 412
	IfUnmodifiedSince *time.Time

	// Part number of the object being read. This is a positive integer between
	// 1 and 10,000. Effectively performs a 'ranged' GET request for the part specified.
	// Useful for downloading just a part of an object.
	PartNumber *int64

	// Downloads the specified range bytes of an object. For more information about
	Range *string

	// ** location: querystring

	// Sets the Cache-Control header of the response.
	ResponseCacheControl *string

	// Sets the Content-Disposition header of the response
	ResponseContentDisposition *string

	// Sets the Content-Encoding header of the response.
	ResponseContentEncoding *string

	// Sets the Content-Language header of the response.
	ResponseContentLanguage *string

	// Sets the Content-Type header of the response.
	ResponseContentType *string

	// Sets the Expires header of the response.
	ResponseExpires *time.Time

	// VersionId used to reference a specific version of the object.
	VersionId *string
}
