package blogconsts

const (
	ErrorCreatingBlog           = "error creating blog"
	ErrorGettingBlog            = "error getting blog"
	ErrorGettingBlogs           = "error getting blogs"
	ErrorUpdatingBlog           = "error updating blog"
	ErrorDeletingBlog           = "error deleting blog"
	ErrorAddingRemovingReaction = "error adding or removing reaction"
	ErrorAddingComment          = "error adding comment"
	ErrorGettingComments        = "error getting comments"
	ErrorDeletingComment        = "error deleting comment"
)

const (
	BlogIDRequired     = "required blog id"
	ReactionIDRequired = "required reaction id"
	InvalidReactionID  = "invalid reaction id"
	InvalidCommentID   = "invalid comment id"
)

const (
	BlogCreatedSuccessfully      = "blog created successfully"
	BlogFetchSuccessfully        = "blog fetched successfully"
	BlogsFetchSuccessfully       = "blogs fetched successfully"
	BlogsFetchSuccessfullyOfUser = "blogs fetched successfully of user"
	BlogUpdatedSuccessfully      = "blog updated successfully"
	BlogDeletedSuccessfully      = "blog deleted successfully"
	ReactionAddedSuccessfully    = "reaction added successfully"
	CommentAddedSuccessfully     = "comment added successfully"
	CommentsFetchSuccessfully    = "comments fetched successfully"
	CommentDeletedSuccessfully   = "comment deleted successfully"
)

const (
	BlogID     = "blog_id"
	BlogIDs    = "blog_ids"
	ReactionID = "reaction_id"
	CommentID  = "comment_id"
	CommentIDs = "comment_ids"
)

const (
	YouAreNotAuthorizedToDeleteThisBlog    = "you are not authorized to delete this blog"
	YouAreNotAuthorizedToGetComments       = "you are not authorized to get comments"
	YouAreNotAuthorizedToDeleteThisComment = "you are not authorized to delete this comment"
)

var ReactionTypes = map[uint64]string{
	1: "like",
	2: "love",
	3: "care",
	4: "haha",
	5: "wow",
	6: "sad",
	7: "angry",
}
