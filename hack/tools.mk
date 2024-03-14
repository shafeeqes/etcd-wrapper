# SPDX-FileCopyrightText: 2024 SAP SE or an SAP affiliate company and Gardener contributors
#
# SPDX-License-Identifier: Apache-2.0

TOOLS_DIR                  := hack/tools
TOOLS_BIN_DIR              := $(TOOLS_DIR)/bin
GOLANGCI_LINT              := $(TOOLS_BIN_DIR)/golangci-lint
GO_ADD_LICENSE             := $(TOOLS_BIN_DIR)/addlicense

# default tool versions
GOLANGCI_LINT_VERSION ?= v1.51.2
GO_ADD_LICENSE_VERSION ?= latest

export TOOLS_BIN_DIR := $(TOOLS_BIN_DIR)
export PATH := $(abspath $(TOOLS_BIN_DIR)):$(PATH)
$(info "TOOLS_BIN_DIR from tools.mk", $(TOOLS_BIN_DIR))
$(info "TOOLS_DIR from tools.mk", $(TOOLS_DIR))
$(info "PATH from tools.mk", $(PATH))

#########################################
# Tools                                 #
#########################################

$(GOLANGCI_LINT): $(call tool_version_file,$(GOLANGCI_LINT),$(GOLANGCI_LINT_VERSION))
	@# CGO_ENABLED has to be set to 1 in order for golangci-lint to be able to load plugins
	@# see https://github.com/golangci/golangci-lint/issues/1276
	GOBIN=$(abspath $(TOOLS_BIN_DIR)) CGO_ENABLED=1 go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VERSION)

$(GO_ADD_LICENSE):
	GOBIN=$(abspath $(TOOLS_BIN_DIR)) go install github.com/google/addlicense@$(GO_ADD_LICENSE_VERSION)