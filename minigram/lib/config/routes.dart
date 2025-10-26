import 'package:flutter/material.dart';
import 'package:minigram/view/splash/splash_screen.dart';

class Routes {
  static String get splashScreen => '/splash_screen';

  static Map<String, WidgetBuilder> get routerApp => {
    splashScreen: (context) => SplashScreen(),
  };
}
