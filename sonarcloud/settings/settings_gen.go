package settings

// AUTOMATICALLY GENERATED, DO NOT EDIT BY HAND!

// ListDefinitionsRequest List settings definitions.<br>Requires 'Browse' permission when a component is specified<br/>To access licensed settings, authentication is required<br/>To access secured settings, one of the following permissions is required: <ul><li>'Execute Analysis'</li><li>'Administer' rights on the specified component</li></ul>
type ListDefinitionsRequest struct {
	Component string `form:"component,omitempty"` // Component key
}

// ListDefinitionsResponse is the response for ListDefinitionsRequest
type ListDefinitionsResponse struct {
	Definitions []struct {
		Category     string   `json:"category,omitempty"`
		DefaultValue string   `json:"defaultValue,omitempty"`
		Description  string   `json:"description,omitempty"`
		Fields       []string `json:"fields,omitempty"`
		Key          string   `json:"key,omitempty"`
		MultiValues  bool     `json:"multiValues,omitempty"`
		Name         string   `json:"name,omitempty"`
		Options      []string `json:"options,omitempty"`
		SubCategory  string   `json:"subCategory,omitempty"`
		Type         string   `json:"type,omitempty"`
	} `json:"definitions,omitempty"`
}

// ResetRequest Remove a setting value.<br>The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires the permission 'Administer' on the specified component.
type ResetRequest struct {
	Branch       string `form:"branch,omitempty"`       // Branch key
	Component    string `form:"component,omitempty"`    // Component key
	Keys         string `form:"keys,omitempty"`         // Comma-separated list of keys
	Organization string `form:"organization,omitempty"` // Organization key
	PullRequest  string `form:"pullRequest,omitempty"`  // Pull request id
}

// SetRequest Update a setting value.<br>Either 'value' or 'values' must be provided.<br> The settings defined in conf/sonar.properties are read-only and can't be changed.<br/>Requires the permission 'Administer' on the specified component.
type SetRequest struct {
	Component    string `form:"component,omitempty"`    // Component key
	FieldValues  string `form:"fieldValues,omitempty"`  // Setting field values. To set several values, the parameter must be called once for each value.
	Key          string `form:"key,omitempty"`          // Setting key
	Organization string `form:"organization,omitempty"` // Organization key (for the Enterprise plan only)
	Value        string `form:"value,omitempty"`        // Setting value. To reset a value, please use the reset web service.
	Values       string `form:"values,omitempty"`       // Setting multi value. To set several values, the parameter must be called once for each value.
}

// ValuesRequest List settings values.<br>If no value has been set for a setting, then the default value is returned.<br>Both component and organization parameters cannot be used together.<br/>Requires 'Browse' or 'Execute Analysis' permission when a component is specified.<br/>Requires to be member of the organization if one is specified.<br/>To access secured settings, one of the following permissions is required: 'Execute Analysis' or 'Administer' rights on the specified component<br/>The returned attributes are:<ul><li>'key': The key of the setting</li><li>'value': The value of setting</li><li>'inherited': True if the value is being inherited from a parent setting</li><li>'parentValue: The value of the parent setting if the value is not inherited'</li><li>'parentOrigin: The origin of the parentValue (INSTANCE, ORGANIZATION, PROJECT)'</li></ul>
type ValuesRequest struct {
	Component    string `form:"component,omitempty"`    // Component key
	Keys         string `form:"keys,omitempty"`         // List of setting keys
	Organization string `form:"organization,omitempty"` // Organization key
}

// ValuesResponse is the response for ValuesRequest
type ValuesResponse struct {
	Settings []struct {
		Inherited    bool     `json:"inherited,omitempty"`
		Key          string   `json:"key,omitempty"`
		ParentOrigin string   `json:"parentOrigin,omitempty"`
		Value        string   `json:"value,omitempty"`
		Values       []string `json:"values,omitempty"`
		FieldValues  []struct {
			Boolean string `json:"boolean,omitempty"`
			Text    string `json:"text,omitempty"`
		} `json:"fieldValues,omitempty"`
	} `json:"settings,omitempty"`
}
