package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/bubbles/list"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Custom key mapping
type keyMap struct {
	Up        key.Binding
	Down      key.Binding
	Enter     key.Binding
	Quit      key.Binding
	Back      key.Binding
	Toggle    key.Binding
	NextStep  key.Binding
	PrevStep  key.Binding
	Help      key.Binding
	Customize key.Binding
	}


func (k keyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Enter, k.Toggle, k.Back, k.Quit}
}

func (k keyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Enter, k.Toggle},
		{k.Back, k.NextStep, k.PrevStep, k.Quit},
	}
}

var keys = keyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "down"),
	),
	Enter: key.NewBinding(
		key.WithKeys("enter"),
		key.WithHelp("enter", "select"),
	),
	Quit: key.NewBinding(
		key.WithKeys("ctrl+c", "q"),
		key.WithHelp("q", "quit"),
	),
	Back: key.NewBinding(
		key.WithKeys("esc"),
		key.WithHelp("esc", "back"),
	),
	Toggle: key.NewBinding(
		key.WithKeys("space"),
		key.WithHelp("space", "toggle"),
	),
	NextStep: key.NewBinding(
		key.WithKeys("tab", "right"),
		key.WithHelp("tab/→", "next"),
	),
	PrevStep: key.NewBinding(
		key.WithKeys("shift+tab", "left"),
		key.WithHelp("shift+tab/←", "previous"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "help"),
	),
	Customize: key.NewBinding(
		key.WithKeys("c"),
		key.WithHelp("c", "customize"),
	),
}

// Step constants for the wizard flow
const (
	StepDatabase    = 0
	StepORM         = 1
	StepAuth        = 2
	StepDocker      = 3
	StepFeatures    = 4
	StepConfirm     = 5
	StepCustomizing = 6
)

// The main model for the TUI
type model struct {
	projectName  string
	config       *ProjectConfig
	step         int
	dbList       list.Model
	ormList      list.Model
	authList     list.Model
	dockerList   list.Model
	featuresList list.Model
	helpModel    help.Model
	width        int
	height       int
	ready        bool
	goodbye      bool
}

// Initialize the model with a project name
func initialModel(projectName string) model {
	// Initialize default configuration
	config := &ProjectConfig{
		ProjectName:    projectName,
		Database:       "PostgreSQL",
		ORM:            "GORM",
		Auth:           "JWT",
		Docker:         "Full (development + production)",
		Features:       []string{},
		GraphQLLibrary: "gqlgen",
	}

	// Set up list delegates
	delegateKeys := list.DefaultKeyMap()
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#00FF00"))
	delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#00FF00"))
	
	// Database list
	dbItems := make([]list.Item, len(DatabaseOptions))
	for i, db := range DatabaseOptions {
		dbItems[i] = item{title: db, desc: getDBDescription(db)}
	}
	dbList := list.New(dbItems, delegate, 0, 0)
	dbList.Title = "Choose a Database"
	dbList.SetShowStatusBar(false)
	dbList.SetFilteringEnabled(false)
	dbList.SetShowHelp(false)
	dbList.KeyMap = delegateKeys

	// We'll initialize the other lists later in Update since they depend on previous selections
	return model{
		projectName:  projectName,
		config:       config,
		step:         StepDatabase,
		dbList:       dbList,
		helpModel:    help.New(),
		ready:        false,
		goodbye:      false,
	}
}

// Item represents a list item with title and description
type item struct {
	title string
	desc  string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

// Start the BubbleTea program
func StartProjectCreator(projectName string) (*ProjectConfig, error) {
	m := initialModel(projectName)
	p := tea.NewProgram(m, tea.WithAltScreen())
	
	result, err := p.Run()
	if err != nil {
		return nil, err
	}
	
	finalModel := result.(model)
	if finalModel.goodbye {
		return nil, nil
	}
	
	return finalModel.config, nil
}

// Get descriptions for database options
func getDBDescription(db string) string {
	switch db {
	case "PostgreSQL":
		return "A powerful, open source object-relational database system"
	case "MongoDB":
		return "A NoSQL document database with JSON-like documents"
	case "MySQL":
		return "The world's most popular open source relational database"
	case "SQLite":
		return "A lightweight disk-based database that doesn't require a server"
	default:
		return ""
	}
}

// Get descriptions for ORM options
func getORMDescription(orm string) string {
	switch orm {
	case "GORM":
		return "The most popular Go ORM with full-featured support"
	case "Ent":
		return "An entity framework from Facebook with code generation"
	case "SQLC":
		return "Generate type-safe code from SQL, not an ORM but very efficient"
	case "Raw SQL":
		return "Direct database access with no ORM abstraction"
	case "Official Go Driver":
		return "The official MongoDB driver for Go"
	case "mgm":
		return "A MongoDB object-document-mapper for Go"
	case "Raw Driver":
		return "Direct MongoDB driver access with no abstraction"
	default:
		return ""
	}
}

// Get descriptions for auth options
func getAuthDescription(auth string) string {
	switch auth {
	case "JWT":
		return "JSON Web Tokens for stateless authentication"
	case "OAuth":
		return "Delegated authorization framework for third-party access"
	case "None":
		return "No authentication, suitable for internal or public APIs"
	default:
		return ""
	}
}

// Get descriptions for Docker options
func getDockerDescription(docker string) string {
	switch docker {
	case "None":
		return "No Docker configuration"
	case "Development only":
		return "Docker setup for local development only"
	case "Full (development + production)":
		return "Complete Docker configuration for both dev and production"
	default:
		return ""
	}
}

// Get descriptions for feature options
func getFeatureDescription(feature string) string {
	switch feature {
	case "WebSocket Subscriptions":
		return "Real-time updates with GraphQL subscriptions"
	case "Redis Caching":
		return "Performance improvement with Redis caching"
	case "Background Jobs":
		return "Async task processing with background workers"
	case "Metrics & Monitoring":
		return "Built-in observability with Prometheus and OpenTelemetry"
	case "File Upload Support":
		return "Support for file uploads through GraphQL mutations"
	default:
		return ""
	}
}

// The Init function for the Bubbletea model
func (m model) Init() tea.Cmd {
	return nil
}

// The Update function for the Bubbletea model
func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch {
		case key.Matches(msg, keys.Quit):
			m.goodbye = true
			return m, tea.Quit

		case key.Matches(msg, keys.NextStep):
			if m.step < StepConfirm {
				m.step++
				return m, nil
			}

		case key.Matches(msg, keys.PrevStep):
			if m.step > StepDatabase {
				m.step--
				return m, nil
			}

		case key.Matches(msg, keys.Back):
			if m.step > StepDatabase {
				m.step--
				return m, nil
			} else {
				m.goodbye = true
				return m, tea.Quit
			}
		}

	case tea.WindowSizeMsg:
		m.width = msg.Width
		m.height = msg.Height
		m.ready = true

		// Update the size of all lists
		headerHeight := 5
		footerHeight := 3
		verticalMargins := headerHeight + footerHeight

		m.dbList.SetSize(msg.Width, msg.Height-verticalMargins)
		
		// Initialize the other lists now that we know the window size
		if m.ormList.Items() == nil {
			// Initialize ORM list based on selected database
			ormOptions := GetORMOptions(m.config.Database)
			ormItems := make([]list.Item, len(ormOptions))
			for i, orm := range ormOptions {
				ormItems[i] = item{title: orm, desc: getORMDescription(orm)}
			}
			
			delegate := list.NewDefaultDelegate()
			delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#00FF00"))
			delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#00FF00"))
			
			m.ormList = list.New(ormItems, delegate, msg.Width, msg.Height-verticalMargins)
			m.ormList.Title = "Select an ORM/Database Access Layer"
			m.ormList.SetShowStatusBar(false)
			m.ormList.SetFilteringEnabled(false)
			m.ormList.SetShowHelp(false)
		}
		
		if m.authList.Items() == nil {
			// Initialize Auth list
			authItems := make([]list.Item, len(AuthOptions))
			for i, auth := range AuthOptions {
				authItems[i] = item{title: auth, desc: getAuthDescription(auth)}
			}
			
			delegate := list.NewDefaultDelegate()
			delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#00FF00"))
			delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#00FF00"))
			
			m.authList = list.New(authItems, delegate, msg.Width, msg.Height-verticalMargins)
			m.authList.Title = "Choose Authentication Method"
			m.authList.SetShowStatusBar(false)
			m.authList.SetFilteringEnabled(false)
			m.authList.SetShowHelp(false)
		}
		
		if m.dockerList.Items() == nil {
			// Initialize Docker list
			dockerItems := make([]list.Item, len(DockerOptions))
			for i, docker := range DockerOptions {
				dockerItems[i] = item{title: docker, desc: getDockerDescription(docker)}
			}
			
			delegate := list.NewDefaultDelegate()
			delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#00FF00"))
			delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#00FF00"))
			
			m.dockerList = list.New(dockerItems, delegate, msg.Width, msg.Height-verticalMargins)
			m.dockerList.Title = "Docker Configuration"
			m.dockerList.SetShowStatusBar(false)
			m.dockerList.SetFilteringEnabled(false)
			m.dockerList.SetShowHelp(false)
		}
		
		if m.featuresList.Items() == nil {
			// Initialize Features list
			featuresItems := make([]list.Item, len(FeatureOptions))
			for i, feature := range FeatureOptions {
				featuresItems[i] = item{title: feature, desc: getFeatureDescription(feature)}
			}
			
			delegate := list.NewDefaultDelegate()
			delegate.Styles.SelectedTitle = delegate.Styles.SelectedTitle.Foreground(lipgloss.Color("#00FF00"))
			delegate.Styles.SelectedDesc = delegate.Styles.SelectedDesc.Foreground(lipgloss.Color("#00FF00"))
			
			m.featuresList = list.New(featuresItems, delegate, msg.Width, msg.Height-verticalMargins)
			m.featuresList.Title = "Enable Additional Features"
			m.featuresList.SetShowStatusBar(false)
			m.featuresList.SetFilteringEnabled(false)
			m.featuresList.SetShowHelp(false)
		}
	}

	// Handle the different steps
	switch m.step {
	case StepDatabase:
		m.dbList, cmd = m.dbList.Update(msg)
		cmds = append(cmds, cmd)
		
		if selected, ok := msg.(tea.KeyMsg); ok && key.Matches(selected, keys.Enter) {
			idx := m.dbList.Index()
			if idx >= 0 && idx < len(DatabaseOptions) {
				m.config.Database = DatabaseOptions[idx]
				
				// Update ORM options based on selected database
				ormOptions := GetORMOptions(m.config.Database)
				ormItems := make([]list.Item, len(ormOptions))
				for i, orm := range ormOptions {
					ormItems[i] = item{title: orm, desc: getORMDescription(orm)}
				}
				m.ormList.SetItems(ormItems)
				
				// Set default ORM
				m.config.ORM = ormOptions[0]
				
				m.step = StepORM
			}
		}
		
	case StepORM:
		m.ormList, cmd = m.ormList.Update(msg)
		cmds = append(cmds, cmd)
		
		if selected, ok := msg.(tea.KeyMsg); ok && key.Matches(selected, keys.Enter) {
			idx := m.ormList.Index()
			if idx >= 0 && idx < len(GetORMOptions(m.config.Database)) {
				m.config.ORM = GetORMOptions(m.config.Database)[idx]
				m.step = StepAuth
			}
		}
		
	case StepAuth:
		m.authList, cmd = m.authList.Update(msg)
		cmds = append(cmds, cmd)
		
		if selected, ok := msg.(tea.KeyMsg); ok && key.Matches(selected, keys.Enter) {
			idx := m.authList.Index()
			if idx >= 0 && idx < len(AuthOptions) {
				m.config.Auth = AuthOptions[idx]
				m.step = StepDocker
			}
		}
		
	case StepDocker:
		m.dockerList, cmd = m.dockerList.Update(msg)
		cmds = append(cmds, cmd)
		
		if selected, ok := msg.(tea.KeyMsg); ok && key.Matches(selected, keys.Enter) {
			idx := m.dockerList.Index()
			if idx >= 0 && idx < len(DockerOptions) {
				m.config.Docker = DockerOptions[idx]
				m.step = StepFeatures
			}
		}
		
	case StepFeatures:
		m.featuresList, cmd = m.featuresList.Update(msg)
		cmds = append(cmds, cmd)
		
		if selected, ok := msg.(tea.KeyMsg); ok {
			// Toggle feature selection with space
			if key.Matches(selected, keys.Toggle) {
				idx := m.featuresList.Index()
				if idx >= 0 && idx < len(FeatureOptions) {
					feature := FeatureOptions[idx]
					found := false
					for i, f := range m.config.Features {
						if f == feature {
							// Remove the feature if already selected
							m.config.Features = append(m.config.Features[:i], m.config.Features[i+1:]...)
							found = true
							break
						}
					}
					if !found {
						// Add the feature if not selected
						m.config.Features = append(m.config.Features, feature)
					}
				}
			} else if key.Matches(selected, keys.Enter) {
				m.step = StepConfirm
			}
		}
		
	case StepConfirm:
		if selected, ok := msg.(tea.KeyMsg); ok && key.Matches(selected, keys.Enter) {
			return m, tea.Quit
		}
	}

	return m, tea.Batch(cmds...)
}

// Helper function to get navigation help based on the current step
func (m model) navigationHelp() string {
	helpView := "\n"
	
	if m.step > StepDatabase {
		helpView += "← Back "
	}
	
	if m.step < StepConfirm {
		helpView += "→ Next "
	}
	
	switch m.step {
	case StepFeatures:
		helpView += "Space Toggle "
	case StepConfirm:
		helpView += "Enter Create Project "
	default:
		helpView += "Enter Select "
	}
	
	helpView += "ESC Back q Quit"
	
	// Style the help text
	helpStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#626262"))
	return helpStyle.Render(helpView)
}

// The View function for the Bubbletea model
func (m model) View() string {
	if !m.ready {
		return "Initializing..."
	}

	if m.goodbye {
		return "Project creation cancelled. Goodbye!\n"
	}

	// Style for header
	headerStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#00FF00")).
		BorderStyle(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#00FF00")).
		Padding(0, 1)

	s := headerStyle.Render(fmt.Sprintf("Creating project: %s", m.projectName)) + "\n\n"

	switch m.step {
	case StepDatabase:
		s += m.dbList.View()
		
	case StepORM:
		s += m.ormList.View()
		
	case StepAuth:
		s += m.authList.View()
		
	case StepDocker:
		s += m.dockerList.View()
		
	case StepFeatures:
		s += "Use space to select/deselect features, enter to continue\n\n"
		s += m.featuresList.View()
		s += "\nSelected features: " + strings.Join(m.config.Features, ", ")
		
	case StepConfirm:
		// Style for the summary
		titleStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Bold(true)
		labelStyle := lipgloss.NewStyle().Foreground(lipgloss.Color("#FFFFFF")).Bold(true)
		
		s += titleStyle.Render("📋 Project Summary") + "\n\n"
		s += labelStyle.Render("Project Name: ") + m.config.ProjectName + "\n"
		s += labelStyle.Render("Database: ") + m.config.Database + "\n"
		s += labelStyle.Render("ORM/DB Access: ") + m.config.ORM + "\n"
		s += labelStyle.Render("Authentication: ") + m.config.Auth + "\n"
		s += labelStyle.Render("Docker: ") + m.config.Docker + "\n"
		
		if len(m.config.Features) > 0 {
			s += labelStyle.Render("Features:") + "\n"
			for _, feature := range m.config.Features {
				s += "  - " + feature + "\n"
			}
		} else {
			s += labelStyle.Render("Features: ") + "None\n"
		}

	}
			// Add navigation help at the bottom
	s += m.navigationHelp()
	
	// Center everything in the available space
	return lipgloss.Place(
		m.width,
		m.height,
		lipgloss.Center,
		lipgloss.Center,
		s,
	)
}