package project_branches

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// DeleteRequest Delete a non-main branch of a project.<br/>Requires 'Administer' rights on the specified project.
type DeleteRequest struct {
	Branch  string `form:"branch,omitempty"`  // Name of the branch
	Project string `form:"project,omitempty"` // Project key
}

// ListRequest List the branches of a project.<br/>The statistics are the overall counts on long branches, and the count of issues detected on the changed code on short branches, and are only provided if the project parameter is specified.<br/>If the project parameter is specified, requires the user to have 'Browse' or 'Execute analysis' rights on that project. Otherwise, only returns branches from projects on which this user has 'Browse' or 'Execute analysis' rights.
type ListRequest struct {
	BranchIds string `form:"branchIds,omitempty"` // List of up to 50 branch IDs - required unless project key is provided
	Project   string `form:"project,omitempty"`   // Project key - required unless branchIds is provided
}

// ListResponse is the response for ListRequest
type ListResponse struct {
	Branches []struct {
		AnalysisDate string `json:"analysisDate,omitempty"`
		BranchId     string `json:"branchId,omitempty"`
		Commit       struct {
			Sha string `json:"sha,omitempty"`
		} `json:"commit,omitempty"`
		IsMain      bool   `json:"isMain,omitempty"`
		MergeBranch string `json:"mergeBranch,omitempty"`
		Name        string `json:"name,omitempty"`
		Status      struct {
			Bugs              float64 `json:"bugs,omitempty"`
			CodeSmells        float64 `json:"codeSmells,omitempty"`
			QualityGateStatus string  `json:"qualityGateStatus,omitempty"`
			Vulnerabilities   float64 `json:"vulnerabilities,omitempty"`
		} `json:"status,omitempty"`
		Type string `json:"type,omitempty"`
	} `json:"branches,omitempty"`
}

// RenameRequest Rename the main branch of a project.<br/>Requires 'Administer' permission on the specified project.
type RenameRequest struct {
	Name    string `form:"name,omitempty"`    // New name of the main branch
	Project string `form:"project,omitempty"` // Project key
}
