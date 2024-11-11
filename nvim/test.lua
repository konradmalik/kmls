---@type vim.lsp.ClientConfig
local config = {
    name = "kmls",
    cmd = { "/home/konrad/Code/github.com/konradmalik/kmls/bin/kmls" },
}

vim.api.nvim_create_autocmd("FileType", {
    pattern = "markdown",
    callback = function()
        local lsp = require("pde.lsp")
        lsp.start(config, { bufnr = 0 })
    end,
})
