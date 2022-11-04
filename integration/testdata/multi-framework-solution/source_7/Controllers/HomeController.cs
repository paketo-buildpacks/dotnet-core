using System.Diagnostics;
using core3_1_dependency;
using Microsoft.AspNetCore.Mvc;
using Net6_Dependency;
using source_7_app.Models;

namespace source_7_app.Controllers;

public class HomeController : Controller
{
    private readonly ILogger<HomeController> _logger;

    object[] objects = new object[] { new Dependent_Framework6(), new Dependent_Core31() };

    public HomeController(ILogger<HomeController> logger)
    {
        _logger = logger;
        _logger.LogDebug($"Total Dependencies: {objects.Length}");
    }

    public IActionResult Index()
    {
        return View();
    }

    public IActionResult Privacy()
    {
        return View();
    }

    [ResponseCache(Duration = 0, Location = ResponseCacheLocation.None, NoStore = true)]
    public IActionResult Error()
    {
        return View(new ErrorViewModel { RequestId = Activity.Current?.Id ?? HttpContext.TraceIdentifier });
    }
}
